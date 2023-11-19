// Code generated by MockGen. DO NOT EDIT.
// Source: transaction-service-client/repository.go
//
// Generated by this command:
//
//	mockgen -source=transaction-service-client/repository.go -package=transaction_service_client_mocks -destination=_mocks/transaction-service-client/repository.go
//
// Package transaction_service_client_mocks is a generated GoMock package.
package transaction_service_client_mocks

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

// CheckChapterBought mocks base method.
func (m *MockRepository) CheckChapterBought(userId, chapterId uint64) (*shared_types.CheckChapterBoughtResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChapterBought", userId, chapterId)
	ret0, _ := ret[0].(*shared_types.CheckChapterBoughtResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckChapterBought indicates an expected call of CheckChapterBought.
func (mr *MockRepositoryMockRecorder) CheckChapterBought(userId, chapterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChapterBought", reflect.TypeOf((*MockRepository)(nil).CheckChapterBought), userId, chapterId)
}
