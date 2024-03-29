package chapters_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	books_controller "github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/books/controller"
	books_model "github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/books/model"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/chapters/model"
	chapters_repository "github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/chapters/repository"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/service"
	transaction_service_client "github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/transaction-service-client"
	auth_middleware "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/auth-middleware"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/router"
	shared_types "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/shared-types"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/utils"
	"golang.org/x/sync/singleflight"
)

type chapterContext string

const (
	middleWareChapter chapterContext = "chapter"
)

type DefaultController struct {
	chapterRepository        chapters_repository.Repository
	transactionServiceClient transaction_service_client.Repository
	service                  service.Service
	g                        *singleflight.Group
}

func NewDefaultController(
	chapterRepository chapters_repository.Repository,
	service service.Service,
	transactionServiceClient transaction_service_client.Repository,
) *DefaultController {
	g := &singleflight.Group{}
	return &DefaultController{chapterRepository, transactionServiceClient, service, g}
}
func (ctrl *DefaultController) GetChaptersForBook(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth_middleware.AuthenticatedUserId).(uint64)
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)

	newChapters, err, _ := ctrl.g.Do(fmt.Sprintf("chapters-%d", book.ID), func() (interface{}, error) {
		return ctrl.chapterRepository.FindAllPreviewsByBookId(book.ID)
	})

	if err != nil {
		log.Println("ERROR [GetChaptersForBook - FindAllPreviewsByBookId]: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	chapters := newChapters.([]*model.ChapterPreview)

	if userId != book.AuthorID {
		chapters = utils.Filter(chapters, func(chapter *model.ChapterPreview) bool { return chapter.Status == model.Published })
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chapters)
}

type createChapterRequest struct {
	Name    string  `json:"name"`
	Price   *uint64 `json:"price"`
	Content string  `json:"content"`
}

func (r createChapterRequest) isValid() bool {
	return r.Name != "" && r.Price != nil && r.Content != ""
}

func (ctrl *DefaultController) PostChapter(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth_middleware.AuthenticatedUserId).(uint64)
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)

	if userId != book.AuthorID {
		log.Println("ERROR [PostChapter - userId != book.AuthorID]: ", "You are not the owner of the book")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var request createChapterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("ERROR [PostChapter - Decode createChapterRequest]: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		log.Println("ERROR [PostChapter - Validate createChapterRequest]: ", "Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ctrl.chapterRepository.Create([]*model.Chapter{{
		BookID:  book.ID,
		Name:    request.Name,
		Price:   *request.Price,
		Content: request.Content,
	}}); err != nil {
		log.Println("ERROR [PostChapter - Create]: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) GetChapterForBook(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth_middleware.AuthenticatedUserId).(uint64)
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)
	chapter := r.Context().Value(middleWareChapter).(*model.Chapter)

	if userId == book.AuthorID {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chapter)
		return
	}

	err := ctrl.transactionServiceClient.CheckChapterBought(userId, chapter.ID, chapter.BookID)
	if err != nil {
		log.Println("ERROR [GetChapterForBook - CheckChapterBought]: ", err.Error())
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chapter)
}

type updateChapterRequest struct {
	Name    string        `json:"name"`
	Price   *uint64       `json:"price"`
	Content string        `json:"content"`
	Status  *model.Status `json:"status"`
}

func (ctrl *DefaultController) PatchChapter(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth_middleware.AuthenticatedUserId).(uint64)
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)
	chapter := r.Context().Value(middleWareChapter).(*model.Chapter)

	if userId != book.AuthorID {
		log.Println("ERROR [PatchChapter - userId != book.AuthorID]: ", "You are not the owner of the book")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var request updateChapterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("ERROR [PatchChapter - Decode updateChapterRequest]: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var patchChapter model.ChapterPatch
	if request.Name != "" {
		patchChapter.Name = &request.Name
	}
	if request.Content != "" {
		patchChapter.Content = &request.Content
	}
	if request.Price != nil {
		patchChapter.Price = request.Price
	}
	if request.Status != nil {
		newstatus := *request.Status
		if newstatus > model.Published {
			log.Println("ERROR [DeleteChapter - userId != book.AuthorID]: ", "You cannot change status to a status greater than published")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newstatus < model.Draft {
			log.Println("ERROR [DeleteChapter - userId != book.AuthorID]: ", "You cannot change status to a status less than draft")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if chapter.Status == model.Published && newstatus == model.Draft {
			log.Println("ERROR [DeleteChapter - userId != book.AuthorID]: ", "You cannot change status from published to draft")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if chapter.Status == model.Draft && newstatus == model.Published {
			patchChapter.Status = request.Status
		}
	}

	if err := ctrl.chapterRepository.Update(chapter.ID, chapter.BookID, &patchChapter); err != nil {
		log.Println("ERROR [PatchChapter - Update]: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) DeleteChapter(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth_middleware.AuthenticatedUserId).(uint64)
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)
	chapter := r.Context().Value(middleWareChapter).(*model.Chapter)

	if userId != book.AuthorID {
		log.Println("ERROR [DeleteChapter - userId != book.AuthorID]: ", "You are not the owner of the book")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if chapter.Status == model.Published {
		log.Println("ERROR [DeleteChapter - chapter.Status == model.Published]: ", "Cannot delete a published chapter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ctrl.chapterRepository.Delete([]*model.Chapter{chapter}); err != nil {
		log.Println("ERROR [DeleteChapter - Delete]: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) LoadChapterMiddleware(w http.ResponseWriter, r *http.Request, next router.Next) {
	book := r.Context().Value(books_controller.MiddleWareBook).(*books_model.Book)
	chapterId := r.Context().Value("chapterid").(string)

	id, err := strconv.ParseUint(chapterId, 10, 64)
	if err != nil {
		log.Println("ERROR [LoadChapterMiddleware - ParseUint]: ", err.Error())
		http.Error(w, "can't parse the chapterId", http.StatusBadRequest)
		return
	}

	newChapter, err, _ := ctrl.g.Do(fmt.Sprintf("chapter-%d", id), func() (interface{}, error) {
		return ctrl.chapterRepository.FindByIdAndBookId(id, book.ID)
	})
	if err != nil {
		log.Println("ERROR [LoadChapterMiddleware - FindByIdAndBookId]: ", err.Error())
		http.Error(w, "can't find the chapter", http.StatusNotFound)
		return
	}
	chapter := newChapter.(*model.Chapter)

	ctx := context.WithValue(r.Context(), middleWareChapter, chapter)
	next(r.WithContext(ctx))
}

func (ctrl *DefaultController) ValidateChapterId(w http.ResponseWriter, r *http.Request) {
	var request shared_types.ValidateChapterIdRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("ERROR [ValidateChapterId - Decode ValidateChapterIdRequest]: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		log.Println("ERROR [ValidateChapterId - Validate ValidateChapterIdRequest]: ", "Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, statusCode, err := ctrl.service.ValidateChapterId(request.UserId, request.ChapterId, request.BookId)
	if err != nil {
		log.Println("ERROR [ValidateChapterId - Execute ValidateChapterId]: ", err.Error())
		http.Error(w, err.Error(), statusCode.ToHTTPStatusCode())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
