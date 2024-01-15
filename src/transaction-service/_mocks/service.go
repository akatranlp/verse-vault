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

// CheckChapterBought mocks base method.
func (m *MockService) CheckChapterBought(userId, chapterId, bookId uint64) (bool, shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChapterBought", userId, chapterId, bookId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(shared_types.Code)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckChapterBought indicates an expected call of CheckChapterBought.
func (mr *MockServiceMockRecorder) CheckChapterBought(userId, chapterId, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChapterBought", reflect.TypeOf((*MockService)(nil).CheckChapterBought), userId, chapterId, bookId)
}