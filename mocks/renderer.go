// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	configuration "github.com/konstellation-io/kli/internal/services/configuration"
)

// MockRenderer is a mock of Renderer interface.
type MockRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockRendererMockRecorder
}

// MockRendererMockRecorder is the mock recorder for MockRenderer.
type MockRendererMockRecorder struct {
	mock *MockRenderer
}

// NewMockRenderer creates a new mock instance.
func NewMockRenderer(ctrl *gomock.Controller) *MockRenderer {
	mock := &MockRenderer{ctrl: ctrl}
	mock.recorder = &MockRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderer) EXPECT() *MockRendererMockRecorder {
	return m.recorder
}

// RenderServers mocks base method.
func (m *MockRenderer) RenderServers(servers []configuration.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RenderServers", servers)
}

// RenderServers indicates an expected call of RenderServers.
func (mr *MockRendererMockRecorder) RenderServers(servers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderServers", reflect.TypeOf((*MockRenderer)(nil).RenderServers), servers)
}
