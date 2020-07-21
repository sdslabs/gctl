// Code generated by MockGen. DO NOT EDIT.
// Source: authentication.go

// Package testmocks is a generated GoMock package.
package testmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	openapi "github.com/sdslabs/gctl/client"
	http "net/http"
	reflect "reflect"
)

// MockAuthAPIService is a mock of AuthAPIService interface
type MockAuthAPIService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthAPIServiceMockRecorder
}

// MockAuthAPIServiceMockRecorder is the mock recorder for MockAuthAPIService
type MockAuthAPIServiceMockRecorder struct {
	mock *MockAuthAPIService
}

// NewMockAuthAPIService creates a new mock instance
func NewMockAuthAPIService(ctrl *gomock.Controller) *MockAuthAPIService {
	mock := &MockAuthAPIService{ctrl: ctrl}
	mock.recorder = &MockAuthAPIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthAPIService) EXPECT() *MockAuthAPIServiceMockRecorder {
	return m.recorder
}

// Refresh mocks base method
func (m *MockAuthAPIService) Refresh(ctx context.Context, authorization string) (openapi.LoginResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx, authorization)
	ret0, _ := ret[0].(openapi.LoginResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Refresh indicates an expected call of Refresh
func (mr *MockAuthAPIServiceMockRecorder) Refresh(ctx, authorization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockAuthAPIService)(nil).Refresh), ctx, authorization)
}
