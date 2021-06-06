// Code generated by MockGen. DO NOT EDIT.
// Source: facade.go

// Package router is a generated GoMock package.
package router

import (
	reflect "reflect"

	chi "github.com/go-chi/chi/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockFacade is a mock of Facade interface.
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade.
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance.
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// Routing mocks base method.
func (m *MockFacade) Routing(mux *chi.Mux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Routing", mux)
}

// Routing indicates an expected call of Routing.
func (mr *MockFacadeMockRecorder) Routing(mux interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routing", reflect.TypeOf((*MockFacade)(nil).Routing), mux)
}
