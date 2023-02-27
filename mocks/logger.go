// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Interface interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", msg)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), msg)
}

// Error mocks base method.
func (m *MockLogger) Error(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", msg)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), msg)
}

// Info mocks base method.
func (m *MockLogger) Info(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", msg)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), msg)
}

// SetDebugLevel mocks base method.
func (m *MockLogger) SetDebugLevel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDebugLevel")
}

// SetDebugLevel indicates an expected call of SetDebugLevel.
func (mr *MockLoggerMockRecorder) SetDebugLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebugLevel", reflect.TypeOf((*MockLogger)(nil).SetDebugLevel))
}

// Success mocks base method.
func (m *MockLogger) Success(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Success", msg)
}

// Success indicates an expected call of Success.
func (mr *MockLoggerMockRecorder) Success(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Success", reflect.TypeOf((*MockLogger)(nil).Success), msg)
}

// Warn mocks base method.
func (m *MockLogger) Warn(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warn", msg)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), msg)
}
