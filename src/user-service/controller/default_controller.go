package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	auth_middleware "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/auth-middleware"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/crypto"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/router"
	shared_types "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/shared-types"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/utils"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/auth"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/repository"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/service"
	"golang.org/x/sync/singleflight"

	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/model"
)

type contextKey int

const authenticatedUserKey contextKey = 0

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (r *loginRequest) isValid() bool {
	return r.Email != "" && r.Password != ""
}

type DefaultController struct {
	userRepository        repository.Repository
	service               service.Service
	hasher                crypto.Hasher
	accessTokenGenerator  auth.TokenGenerator
	refreshTokenGenerator auth.TokenGenerator
	authIsActive          bool
	g                     *singleflight.Group
}

func NewDefaultController(
	userRepository repository.Repository,
	service service.Service,
	hasher crypto.Hasher,
	accessTokenGenerator auth.TokenGenerator,
	refreshTokenGenerator auth.TokenGenerator,
	authIsActive bool,
) *DefaultController {
	g := &singleflight.Group{}
	return &DefaultController{userRepository, service, hasher, accessTokenGenerator, refreshTokenGenerator, authIsActive, g}
}

func (ctrl *DefaultController) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := ctrl.userRepository.FindByEmail(request.Email)
	if err != nil {
		log.Printf("could not find user by email: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(users) < 1 {
		w.Header().Add("WWW-Authenticate", "Bearer")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if ok := ctrl.hasher.Validate([]byte(request.Password), users[0].Password); !ok {
		w.Header().Add("WWW-Authenticate", "Bearer")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, err := ctrl.accessTokenGenerator.CreateToken(map[string]interface{}{
		"id":            users[0].ID,
		"email":         users[0].Email,
		"token_version": users[0].TokenVersion,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refreshToken, err := ctrl.refreshTokenGenerator.CreateToken(map[string]interface{}{
		"id":            users[0].ID,
		"email":         users[0].Email,
		"token_version": users[0].TokenVersion,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		MaxAge:   int(ctrl.refreshTokenGenerator.GetTokenExpiration().Seconds()),
		HttpOnly: true,
		Path:     "/api/v1/refresh-token",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &newCookie)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(ctrl.accessTokenGenerator.GetTokenExpiration().Seconds()),
	})
}

type registerRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	ProfileName string `json:"profileName"`
}

func (r *registerRequest) isValid() bool {
	return r.Email != "" && r.Password != "" && r.ProfileName != ""
}

func (ctrl *DefaultController) Register(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := ctrl.userRepository.FindByEmail(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(user) > 0 {
		w.WriteHeader(http.StatusConflict)
		return
	}

	hashedPassword, err := ctrl.hasher.Hash([]byte(request.Password))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := ctrl.userRepository.Create([]*model.DbUser{{
		Email:       request.Email,
		Password:    hashedPassword,
		ProfileName: request.ProfileName,
	}}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ctrl *DefaultController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token := ""
	if ctrl.authIsActive {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			log.Println("ERROR [REFRESH_TOKEN - get cookie]: ", err.Error())
			http.Error(w, "There was no cookie in the request!", http.StatusUnauthorized)
			return
		}
		token = cookie.Value
	}

	user, statusCode, err := ctrl.service.ValidateRefreshToken(token)
	if user == nil {
		http.Error(w, err.Error(), statusCode.ToHTTPStatusCode())
		return
	}

	accessToken, err := ctrl.accessTokenGenerator.CreateToken(map[string]interface{}{
		"id":            user.ID,
		"email":         user.Email,
		"token_version": user.TokenVersion,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refreshToken, err := ctrl.refreshTokenGenerator.CreateToken(map[string]interface{}{
		"id":            user.ID,
		"email":         user.Email,
		"token_version": user.TokenVersion,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		MaxAge:   int(ctrl.refreshTokenGenerator.GetTokenExpiration().Seconds()),
		HttpOnly: true,
		Path:     "/api/v1/refresh-token",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &newCookie)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(ctrl.accessTokenGenerator.GetTokenExpiration().Seconds()),
	})
}

func (ctrl *DefaultController) Logout(w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all")

	if all != "" {
		user := r.Context().Value(authenticatedUserKey).(*model.DbUser)
		newTokenVersion := user.TokenVersion + 1
		if err := ctrl.userRepository.Update(user.ID, &model.DbUserPatch{TokenVersion: &newTokenVersion}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	newCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/api/v1/refresh-token",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &newCookie)
	w.WriteHeader(http.StatusOK)
}

func (ctrl *DefaultController) GetUsers(w http.ResponseWriter, _ *http.Request) {
	newUsers, err, _ := ctrl.g.Do("get-users", func() (interface{}, error) {
		return ctrl.userRepository.FindAll()
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users := newUsers.([]*model.DbUser)

	userDto := utils.Map(users, func(user *model.DbUser) model.UserDTO {
		return user.ToDto()
	})

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDto)
}

func (ctrl *DefaultController) GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(authenticatedUserKey).(*model.DbUser)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToDto())
}

func (ctrl *DefaultController) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userid").(string)

	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err, _ := ctrl.g.Do(fmt.Sprintf("user-%d", id), func() (interface{}, error) {
		return ctrl.userRepository.FindById(id)
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user := newUser.(*model.DbUser)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToDto())
}

type putMeRequest struct {
	Password    string `json:"password"`
	ProfileName string `json:"profileName"`
	Balance     *int64 `json:"balance"`
}

func (ctrl *DefaultController) PatchMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(authenticatedUserKey).(*model.DbUser)

	var request putMeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var patchUser model.DbUserPatch

	if request.Password != "" {
		hashedPassword, err := ctrl.hasher.Hash([]byte(request.Password))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		patchUser.Password = &hashedPassword
	}
	if request.ProfileName != "" {
		patchUser.ProfileName = &request.ProfileName
	}
	if request.Balance != nil {
		patchUser.Balance = request.Balance
	}

	if err := ctrl.userRepository.Update(user.ID, &patchUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) DeleteMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(authenticatedUserKey).(*model.DbUser)
	if err := ctrl.userRepository.Delete([]*model.DbUser{user}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type validateTokenRequest struct {
	Token string `json:"token"`
}

func (r *validateTokenRequest) isValid() bool {
	return r.Token != ""
}

func (ctrl *DefaultController) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var request validateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("ERROR [VALIDATE_TOKEN]: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		log.Println("ERROR [VALIDATE_TOKEN]: ", "is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, statusCode, err := ctrl.service.ValidateAccessToken(request.Token)
	if user == nil {
		http.Error(w, err.Error(), statusCode.ToHTTPStatusCode())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth_middleware.VerifyTokenResponse{
		Success: true,
		UserId:  user.ID,
	})
}

func (ctrl *DefaultController) MoveUserAmount(w http.ResponseWriter, r *http.Request) {
	var request shared_types.MoveBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statusCode, err := ctrl.service.MoveUserAmount(request.UserId, request.ReceivingUserId, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), statusCode.ToHTTPStatusCode())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shared_types.MoveBalanceResponse{Success: true})
}

func (ctrl *DefaultController) AuthenticationMiddleWare(w http.ResponseWriter, r *http.Request, next router.Next) {
	if !ctrl.authIsActive {
		user, err := ctrl.userRepository.FindById(1)
		if err != nil {
			http.Error(w, "The user doesn't exist anymore", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), authenticatedUserKey, user)
		next(r.WithContext(ctx))
		return
	}

	token := r.Header.Get("Authorization")

	after, found := strings.CutPrefix(token, "Bearer ")
	if !found {
		http.Error(w, "There was no Token provided", http.StatusUnauthorized)
		return
	}
	user, statusCode, err := ctrl.service.ValidateAccessToken(after)
	if user == nil {
		http.Error(w, err.Error(), statusCode.ToHTTPStatusCode())
		return
	}

	ctx := context.WithValue(r.Context(), authenticatedUserKey, user)
	next(r.WithContext(ctx))
}
