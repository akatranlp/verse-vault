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
	model "github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/model"
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

// ValidateAccessToken mocks base method.
func (m *MockService) ValidateAccessToken(token string) (*model.DbUser, shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessToken", token)
	ret0, _ := ret[0].(*model.DbUser)
	ret1, _ := ret[1].(shared_types.Code)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateAccessToken indicates an expected call of ValidateAccessToken.
func (mr *MockServiceMockRecorder) ValidateAccessToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessToken", reflect.TypeOf((*MockService)(nil).ValidateAccessToken), token)
}

// ValidateRefreshToken mocks base method.
func (m *MockService) ValidateRefreshToken(token string) (*model.DbUser, shared_types.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", token)
	ret0, _ := ret[0].(*model.DbUser)
	ret1, _ := ret[1].(shared_types.Code)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockServiceMockRecorder) ValidateRefreshToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockService)(nil).ValidateRefreshToken), token)
}
