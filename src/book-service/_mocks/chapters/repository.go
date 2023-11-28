// Code generated by MockGen. DO NOT EDIT.
// Source: chapters/repository.go
//
// Generated by this command:
//
//	mockgen -source=chapters/repository.go -package=chapters_mocks -destination=_mocks/chapters/repository.go
//
// Package chapters_mocks is a generated GoMock package.
package chapters_mocks

import (
	reflect "reflect"

	model "github.com/akatranlp/hsfl-master-ai-cloud-engineering/book-service/chapters/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 []*model.Chapter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 []*model.Chapter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0)
}

// FindAllPreviewsByBookId mocks base method.
func (m *MockRepository) FindAllPreviewsByBookId(bookId uint64) ([]*model.ChapterPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllPreviewsByBookId", bookId)
	ret0, _ := ret[0].([]*model.ChapterPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllPreviewsByBookId indicates an expected call of FindAllPreviewsByBookId.
func (mr *MockRepositoryMockRecorder) FindAllPreviewsByBookId(bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllPreviewsByBookId", reflect.TypeOf((*MockRepository)(nil).FindAllPreviewsByBookId), bookId)
}

// FindByIdAndBookId mocks base method.
func (m *MockRepository) FindByIdAndBookId(id, bookId uint64) (*model.Chapter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdAndBookId", id, bookId)
	ret0, _ := ret[0].(*model.Chapter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdAndBookId indicates an expected call of FindByIdAndBookId.
func (mr *MockRepositoryMockRecorder) FindByIdAndBookId(id, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdAndBookId", reflect.TypeOf((*MockRepository)(nil).FindByIdAndBookId), id, bookId)
}

// Migrate mocks base method.
func (m *MockRepository) Migrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockRepositoryMockRecorder) Migrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockRepository)(nil).Migrate))
}

// Update mocks base method.
func (m *MockRepository) Update(id, bookId uint64, updateChapter *model.ChapterPatch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, bookId, updateChapter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(id, bookId, updateChapter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, bookId, updateChapter)
}

// ValidateChapterId mocks base method.
func (m *MockRepository) ValidateChapterId(id, bookId uint64) (*model.Chapter, *uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateChapterId", id, bookId)
	ret0, _ := ret[0].(*model.Chapter)
	ret1, _ := ret[1].(*uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateChapterId indicates an expected call of ValidateChapterId.
func (mr *MockRepositoryMockRecorder) ValidateChapterId(id, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateChapterId", reflect.TypeOf((*MockRepository)(nil).ValidateChapterId), id, bookId)
}
