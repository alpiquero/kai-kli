// Code generated by MockGen. DO NOT EDIT.
// Source: version.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	version "github.com/konstellation-io/kli/api/kre/version"
	reflect "reflect"
)

// MockVersionInterface is a mock of VersionInterface interface
type MockVersionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVersionInterfaceMockRecorder
}

// MockVersionInterfaceMockRecorder is the mock recorder for MockVersionInterface
type MockVersionInterfaceMockRecorder struct {
	mock *MockVersionInterface
}

// NewMockVersionInterface creates a new mock instance
func NewMockVersionInterface(ctrl *gomock.Controller) *MockVersionInterface {
	mock := &MockVersionInterface{ctrl: ctrl}
	mock.recorder = &MockVersionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVersionInterface) EXPECT() *MockVersionInterfaceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockVersionInterface) List() (version.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(version.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVersionInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVersionInterface)(nil).List))
}

// Start mocks base method
func (m *MockVersionInterface) Start(versionID, comment string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", versionID, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockVersionInterfaceMockRecorder) Start(versionID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockVersionInterface)(nil).Start), versionID, comment)
}

// Stop mocks base method
func (m *MockVersionInterface) Stop(versionID, comment string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", versionID, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockVersionInterfaceMockRecorder) Stop(versionID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockVersionInterface)(nil).Stop), versionID, comment)
}

// Publish mocks base method
func (m *MockVersionInterface) Publish(versionID, comment string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", versionID, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockVersionInterfaceMockRecorder) Publish(versionID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockVersionInterface)(nil).Publish), versionID, comment)
}

// Unpublish mocks base method
func (m *MockVersionInterface) Unpublish(versionID, comment string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpublish", versionID, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unpublish indicates an expected call of Unpublish
func (mr *MockVersionInterfaceMockRecorder) Unpublish(versionID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpublish", reflect.TypeOf((*MockVersionInterface)(nil).Unpublish), versionID, comment)
}

// GetConfig mocks base method
func (m *MockVersionInterface) GetConfig(versionID string) (*version.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", versionID)
	ret0, _ := ret[0].(*version.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockVersionInterfaceMockRecorder) GetConfig(versionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockVersionInterface)(nil).GetConfig), versionID)
}

// UpdateConfig mocks base method
func (m *MockVersionInterface) UpdateConfig(versionID string, configVars []version.ConfigVariableInput) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfig", versionID, configVars)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockVersionInterfaceMockRecorder) UpdateConfig(versionID, configVars interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockVersionInterface)(nil).UpdateConfig), versionID, configVars)
}

// Create mocks base method
func (m *MockVersionInterface) Create(krtFile string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", krtFile)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVersionInterfaceMockRecorder) Create(krtFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVersionInterface)(nil).Create), krtFile)
}
