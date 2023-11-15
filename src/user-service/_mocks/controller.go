// Code generated by MockGen. DO NOT EDIT.
// Source: user/controller.go
//
// Generated by this command:
//
//	mockgen.exe -package=mocks -destination=_mocks/controller.go -source=user/controller.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	router "github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/router"
	gomock "go.uber.org/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// AuthenticationMiddleWare mocks base method.
func (m *MockController) AuthenticationMiddleWare(arg0 http.ResponseWriter, arg1 *http.Request, arg2 router.Next) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AuthenticationMiddleWare", arg0, arg1, arg2)
}

// AuthenticationMiddleWare indicates an expected call of AuthenticationMiddleWare.
func (mr *MockControllerMockRecorder) AuthenticationMiddleWare(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationMiddleWare", reflect.TypeOf((*MockController)(nil).AuthenticationMiddleWare), arg0, arg1, arg2)
}

// DeleteMe mocks base method.
func (m *MockController) DeleteMe(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteMe", arg0, arg1)
}

// DeleteMe indicates an expected call of DeleteMe.
func (mr *MockControllerMockRecorder) DeleteMe(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMe", reflect.TypeOf((*MockController)(nil).DeleteMe), arg0, arg1)
}

// GetMe mocks base method.
func (m *MockController) GetMe(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetMe", arg0, arg1)
}

// GetMe indicates an expected call of GetMe.
func (mr *MockControllerMockRecorder) GetMe(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockController)(nil).GetMe), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockController) GetUser(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", arg0, arg1)
}

// GetUser indicates an expected call of GetUser.
func (mr *MockControllerMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockController)(nil).GetUser), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockController) GetUsers(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", arg0, arg1)
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockControllerMockRecorder) GetUsers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockController)(nil).GetUsers), arg0, arg1)
}

// Login mocks base method.
func (m *MockController) Login(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", arg0, arg1)
}

// Login indicates an expected call of Login.
func (mr *MockControllerMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockController)(nil).Login), arg0, arg1)
}

// MoveUserAmount mocks base method.
func (m *MockController) MoveUserAmount(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MoveUserAmount", arg0, arg1)
}

// MoveUserAmount indicates an expected call of MoveUserAmount.
func (mr *MockControllerMockRecorder) MoveUserAmount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveUserAmount", reflect.TypeOf((*MockController)(nil).MoveUserAmount), arg0, arg1)
}

// PatchMe mocks base method.
func (m *MockController) PatchMe(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PatchMe", arg0, arg1)
}

// PatchMe indicates an expected call of PatchMe.
func (mr *MockControllerMockRecorder) PatchMe(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMe", reflect.TypeOf((*MockController)(nil).PatchMe), arg0, arg1)
}

// Register mocks base method.
func (m *MockController) Register(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0, arg1)
}

// Register indicates an expected call of Register.
func (mr *MockControllerMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockController)(nil).Register), arg0, arg1)
}

// ValidateToken mocks base method.
func (m *MockController) ValidateToken(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ValidateToken", arg0, arg1)
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockControllerMockRecorder) ValidateToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockController)(nil).ValidateToken), arg0, arg1)
}
