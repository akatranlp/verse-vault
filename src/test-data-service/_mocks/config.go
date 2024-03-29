// Code generated by MockGen. DO NOT EDIT.
// Source: config/config.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=_mocks/config.go -source=config/config.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// GetSqlString mocks base method.
func (m *MockConfig) GetSqlString() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSqlString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSqlString indicates an expected call of GetSqlString.
func (mr *MockConfigMockRecorder) GetSqlString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSqlString", reflect.TypeOf((*MockConfig)(nil).GetSqlString))
}

// GetUserPassword mocks base method.
func (m *MockConfig) GetUserPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserPassword indicates an expected call of GetUserPassword.
func (mr *MockConfigMockRecorder) GetUserPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockConfig)(nil).GetUserPassword))
}
