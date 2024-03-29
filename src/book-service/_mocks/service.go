// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go
//
// Generated by this command:
//
//	mockgen -source=service/service.go -package=mocks -destination=_mocks/service.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	shared_types "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/shared-types"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ValidateChapterId mocks base method.
func (m *MockService) ValidateChapterId(userId, chapterId, bookId uint64) (*shared_types.ValidateChapterIdResponse, shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateChapterId", userId, chapterId, bookId)
	ret0, _ := ret[0].(*shared_types.ValidateChapterIdResponse)
	ret1, _ := ret[1].(shared_types.Code)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateChapterId indicates an expected call of ValidateChapterId.
func (mr *MockServiceMockRecorder) ValidateChapterId(userId, chapterId, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateChapterId", reflect.TypeOf((*MockService)(nil).ValidateChapterId), userId, chapterId, bookId)
}
