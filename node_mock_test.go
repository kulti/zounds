// Code generated by MockGen. DO NOT EDIT.
// Source: world.go

// Package zounds_test is a generated GoMock package.
package zounds_test

import (
	gomock "github.com/golang/mock/gomock"
	image "image"
	reflect "reflect"
)

// MockStaticNode is a mock of StaticNode interface
type MockStaticNode struct {
	ctrl     *gomock.Controller
	recorder *MockStaticNodeMockRecorder
}

// MockStaticNodeMockRecorder is the mock recorder for MockStaticNode
type MockStaticNodeMockRecorder struct {
	mock *MockStaticNode
}

// NewMockStaticNode creates a new mock instance
func NewMockStaticNode(ctrl *gomock.Controller) *MockStaticNode {
	mock := &MockStaticNode{ctrl: ctrl}
	mock.recorder = &MockStaticNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStaticNode) EXPECT() *MockStaticNodeMockRecorder {
	return m.recorder
}

// Bounds mocks base method
func (m *MockStaticNode) Bounds() image.Rectangle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bounds")
	ret0, _ := ret[0].(image.Rectangle)
	return ret0
}

// Bounds indicates an expected call of Bounds
func (mr *MockStaticNodeMockRecorder) Bounds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bounds", reflect.TypeOf((*MockStaticNode)(nil).Bounds))
}

// Draw mocks base method
func (m *MockStaticNode) Draw() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Draw")
}

// Draw indicates an expected call of Draw
func (mr *MockStaticNodeMockRecorder) Draw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockStaticNode)(nil).Draw))
}
