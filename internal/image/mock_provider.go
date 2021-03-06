// Code generated by MockGen. DO NOT EDIT.
// Source: internal/image/provider.go

// Package image is a generated GoMock package.
package image

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetImage mocks base method
func (m *MockProvider) GetImage(x, y int) (*Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", x, y)
	ret0, _ := ret[0].(*Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage
func (mr *MockProviderMockRecorder) GetImage(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockProvider)(nil).GetImage), x, y)
}
