// Code generated by MockGen. DO NOT EDIT.
// Source: balancer/strategy/strategy.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=_mocks/strategy.go -source=balancer/strategy/strategy.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	target "github.com/akatranlp/hsfl-master-ai-cloud-engineering/load-balancer/balancer/target"
	gomock "go.uber.org/mock/gomock"
)

// MockStrategy is a mock of Strategy interface.
type MockStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockStrategyMockRecorder
}

// MockStrategyMockRecorder is the mock recorder for MockStrategy.
type MockStrategyMockRecorder struct {
	mock *MockStrategy
}

// NewMockStrategy creates a new mock instance.
func NewMockStrategy(ctrl *gomock.Controller) *MockStrategy {
	mock := &MockStrategy{ctrl: ctrl}
	mock.recorder = &MockStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStrategy) EXPECT() *MockStrategyMockRecorder {
	return m.recorder
}

// NextTarget mocks base method.
func (m *MockStrategy) NextTarget(arg0 *http.Request) *target.Target {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextTarget", arg0)
	ret0, _ := ret[0].(*target.Target)
	return ret0
}

// NextTarget indicates an expected call of NextTarget.
func (mr *MockStrategyMockRecorder) NextTarget(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextTarget", reflect.TypeOf((*MockStrategy)(nil).NextTarget), arg0)
}

// SetTargets mocks base method.
func (m *MockStrategy) SetTargets(arg0 []*target.Target) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTargets", arg0)
}

// SetTargets indicates an expected call of SetTargets.
func (mr *MockStrategyMockRecorder) SetTargets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTargets", reflect.TypeOf((*MockStrategy)(nil).SetTargets), arg0)
}
