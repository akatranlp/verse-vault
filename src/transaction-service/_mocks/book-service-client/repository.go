// Code generated by MockGen. DO NOT EDIT.
// Source: book-service-client/repository.go
//
// Generated by this command:
//
//	mockgen -source=book-service-client/repository.go -package=book_service_client_mocks -destination=_mocks/book-service-client/repository.go
//
// Package book_service_client_mocks is a generated GoMock package.
package book_service_client_mocks

import (
	reflect "reflect"

	shared_types "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/shared-types"
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

// ValidateChapterId mocks base method.
func (m *MockRepository) ValidateChapterId(userId, chapterId, bookId uint64) (*shared_types.ValidateChapterIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateChapterId", userId, chapterId, bookId)
	ret0, _ := ret[0].(*shared_types.ValidateChapterIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateChapterId indicates an expected call of ValidateChapterId.
func (mr *MockRepositoryMockRecorder) ValidateChapterId(userId, chapterId, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateChapterId", reflect.TypeOf((*MockRepository)(nil).ValidateChapterId), userId, chapterId, bookId)
}
