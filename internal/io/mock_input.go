// Code generated by MockGen. DO NOT EDIT.
// Source: input.go
//
// Generated by this command:
//
//	mockgen -source=input.go -package=io -destination=mock_input.go
//

// Package io is a generated GoMock package.
package io

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIInputManager is a mock of IInputManager interface.
type MockIInputManager struct {
	ctrl     *gomock.Controller
	recorder *MockIInputManagerMockRecorder
	isgomock struct{}
}

// MockIInputManagerMockRecorder is the mock recorder for MockIInputManager.
type MockIInputManagerMockRecorder struct {
	mock *MockIInputManager
}

// NewMockIInputManager creates a new mock instance.
func NewMockIInputManager(ctrl *gomock.Controller) *MockIInputManager {
	mock := &MockIInputManager{ctrl: ctrl}
	mock.recorder = &MockIInputManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInputManager) EXPECT() *MockIInputManagerMockRecorder {
	return m.recorder
}

// GetCheckboxes mocks base method.
func (m *MockIInputManager) GetCheckboxes(headers, opts []string) ([]string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckboxes", headers, opts)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCheckboxes indicates an expected call of GetCheckboxes.
func (mr *MockIInputManagerMockRecorder) GetCheckboxes(headers, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckboxes", reflect.TypeOf((*MockIInputManager)(nil).GetCheckboxes), headers, opts)
}

// GetYesNo mocks base method.
func (m *MockIInputManager) GetYesNo(label string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYesNo", label)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetYesNo indicates an expected call of GetYesNo.
func (mr *MockIInputManagerMockRecorder) GetYesNo(label any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYesNo", reflect.TypeOf((*MockIInputManager)(nil).GetYesNo), label)
}

// InputKeywordForFilter mocks base method.
func (m *MockIInputManager) InputKeywordForFilter(label string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputKeywordForFilter", label)
	ret0, _ := ret[0].(string)
	return ret0
}

// InputKeywordForFilter indicates an expected call of InputKeywordForFilter.
func (mr *MockIInputManagerMockRecorder) InputKeywordForFilter(label any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputKeywordForFilter", reflect.TypeOf((*MockIInputManager)(nil).InputKeywordForFilter), label)
}
