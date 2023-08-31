// Code generated by MockGen. DO NOT EDIT.
// Source: process_registry.go

// Package mocks is a generated GoMock package.
package mocks

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	processregistry "github.com/konstellation-io/kli/api/processregistry"
	configuration "github.com/konstellation-io/kli/internal/services/configuration"
)

// MockAPIClient is a mock of APIClient interface.
type MockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIClientMockRecorder
}

// MockAPIClientMockRecorder is the mock recorder for MockAPIClient.
type MockAPIClientMockRecorder struct {
	mock *MockAPIClient
}

// NewMockAPIClient creates a new mock instance.
func NewMockAPIClient(ctrl *gomock.Controller) *MockAPIClient {
	mock := &MockAPIClient{ctrl: ctrl}
	mock.recorder = &MockAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIClient) EXPECT() *MockAPIClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockAPIClient) List(server *configuration.Server, productID, processType string) ([]processregistry.RegisteredProcess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", server, productID, processType)
	ret0, _ := ret[0].([]processregistry.RegisteredProcess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAPIClientMockRecorder) List(server, productID, processType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAPIClient)(nil).List), server, productID, processType)
}

// Register mocks base method.
func (m *MockAPIClient) Register(server *configuration.Server, processFile *os.File, productID, processID, processType, version string) (*processregistry.RegisteredProcess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", server, processFile, productID, processID, processType, version)
	ret0, _ := ret[0].(*processregistry.RegisteredProcess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAPIClientMockRecorder) Register(server, processFile, productID, processID, processType, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAPIClient)(nil).Register), server, processFile, productID, processID, processType, version)
}
