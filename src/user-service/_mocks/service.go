// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=_mocks/service.go -source=service/service.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	shared_types "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/shared-types"
	model "github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/user/model"
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

// MoveUserAmount mocks base method.
func (m *MockService) MoveUserAmount(payingUserId, receivingUserId uint64, amount int64) (shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveUserAmount", payingUserId, receivingUserId, amount)
	ret0, _ := ret[0].(shared_types.Code)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoveUserAmount indicates an expected call of MoveUserAmount.
func (mr *MockServiceMockRecorder) MoveUserAmount(payingUserId, receivingUserId, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveUserAmount", reflect.TypeOf((*MockService)(nil).MoveUserAmount), payingUserId, receivingUserId, amount)
}

// ValidateToken mocks base method.
func (m *MockService) ValidateToken(token string) (*model.DbUser, shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(*model.DbUser)
	ret1, _ := ret[1].(shared_types.Code)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockServiceMockRecorder) ValidateToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockService)(nil).ValidateToken), token)
}