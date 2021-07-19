// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/dnaRepositoryInterface.go

// Package repositories is a generated GoMock package.
package repositories

import (
	models "mutantDetector/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDnaRepositoryInterface is a mock of DnaRepositoryInterface interface.
type MockDnaRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDnaRepositoryInterfaceMockRecorder
}

// MockDnaRepositoryInterfaceMockRecorder is the mock recorder for MockDnaRepositoryInterface.
type MockDnaRepositoryInterfaceMockRecorder struct {
	mock *MockDnaRepositoryInterface
}

// NewMockDnaRepositoryInterface creates a new mock instance.
func NewMockDnaRepositoryInterface(ctrl *gomock.Controller) *MockDnaRepositoryInterface {
	mock := &MockDnaRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockDnaRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDnaRepositoryInterface) EXPECT() *MockDnaRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDnaRepositoryInterface) Create(dna *models.Dna) (*models.Dna, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", dna)
	ret0, _ := ret[0].(*models.Dna)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDnaRepositoryInterfaceMockRecorder) Create(dna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDnaRepositoryInterface)(nil).Create), dna)
}

// FindByHash mocks base method.
func (m *MockDnaRepositoryInterface) FindByHash(dna *models.Dna) (*models.Dna, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByHash", dna)
	ret0, _ := ret[0].(*models.Dna)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByHash indicates an expected call of FindByHash.
func (mr *MockDnaRepositoryInterfaceMockRecorder) FindByHash(dna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByHash", reflect.TypeOf((*MockDnaRepositoryInterface)(nil).FindByHash), dna)
}