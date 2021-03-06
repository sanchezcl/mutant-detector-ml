// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/dnaStatsRepositoryInterface.go

// Package repositories is a generated GoMock package.
package repositories

import (
	models "mutantDetector/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDnaStatsRepositoryInterface is a mock of DnaStatsRepositoryInterface interface.
type MockDnaStatsRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDnaStatsRepositoryInterfaceMockRecorder
}

// MockDnaStatsRepositoryInterfaceMockRecorder is the mock recorder for MockDnaStatsRepositoryInterface.
type MockDnaStatsRepositoryInterfaceMockRecorder struct {
	mock *MockDnaStatsRepositoryInterface
}

// NewMockDnaStatsRepositoryInterface creates a new mock instance.
func NewMockDnaStatsRepositoryInterface(ctrl *gomock.Controller) *MockDnaStatsRepositoryInterface {
	mock := &MockDnaStatsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockDnaStatsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDnaStatsRepositoryInterface) EXPECT() *MockDnaStatsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetStats mocks base method.
func (m *MockDnaStatsRepositoryInterface) GetStats() (*models.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(*models.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockDnaStatsRepositoryInterfaceMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockDnaStatsRepositoryInterface)(nil).GetStats))
}
