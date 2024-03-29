// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/repository.go -package=mocks -destination=_mocks/repository.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/akatranlp/hsfl-master-ai-cloud-engineering/transaction-service/model"
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
func (m *MockRepository) Create(arg0 []*model.Transaction) error {
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

// FindAllForReceivingUserId mocks base method.
func (m *MockRepository) FindAllForReceivingUserId(userId uint64) ([]*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllForReceivingUserId", userId)
	ret0, _ := ret[0].([]*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllForReceivingUserId indicates an expected call of FindAllForReceivingUserId.
func (mr *MockRepositoryMockRecorder) FindAllForReceivingUserId(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllForReceivingUserId", reflect.TypeOf((*MockRepository)(nil).FindAllForReceivingUserId), userId)
}

// FindAllForUserId mocks base method.
func (m *MockRepository) FindAllForUserId(userId uint64) ([]*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllForUserId", userId)
	ret0, _ := ret[0].([]*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllForUserId indicates an expected call of FindAllForUserId.
func (mr *MockRepositoryMockRecorder) FindAllForUserId(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllForUserId", reflect.TypeOf((*MockRepository)(nil).FindAllForUserId), userId)
}

// FindForUserIdAndChapterId mocks base method.
func (m *MockRepository) FindForUserIdAndChapterId(userId, chapterId, bookId uint64) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindForUserIdAndChapterId", userId, chapterId, bookId)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindForUserIdAndChapterId indicates an expected call of FindForUserIdAndChapterId.
func (mr *MockRepositoryMockRecorder) FindForUserIdAndChapterId(userId, chapterId, bookId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindForUserIdAndChapterId", reflect.TypeOf((*MockRepository)(nil).FindForUserIdAndChapterId), userId, chapterId, bookId)
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
