// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: UnknownPacketHandler)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUnknownPacketHandler is a mock of UnknownPacketHandler interface.
type MockUnknownPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUnknownPacketHandlerMockRecorder
}

// MockUnknownPacketHandlerMockRecorder is the mock recorder for MockUnknownPacketHandler.
type MockUnknownPacketHandlerMockRecorder struct {
	mock *MockUnknownPacketHandler
}

// NewMockUnknownPacketHandler creates a new mock instance.
func NewMockUnknownPacketHandler(ctrl *gomock.Controller) *MockUnknownPacketHandler {
	mock := &MockUnknownPacketHandler{ctrl: ctrl}
	mock.recorder = &MockUnknownPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnknownPacketHandler) EXPECT() *MockUnknownPacketHandlerMockRecorder {
	return m.recorder
}

// handlePacket mocks base method.
func (m *MockUnknownPacketHandler) handlePacket(arg0 receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockUnknownPacketHandlerMockRecorder) handlePacket(arg0 interface{}) *UnknownPacketHandlerhandlePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockUnknownPacketHandler)(nil).handlePacket), arg0)
	return &UnknownPacketHandlerhandlePacketCall{Call: call}
}

// UnknownPacketHandlerhandlePacketCall wrap *gomock.Call
type UnknownPacketHandlerhandlePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnknownPacketHandlerhandlePacketCall) Return() *UnknownPacketHandlerhandlePacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnknownPacketHandlerhandlePacketCall) Do(f func(receivedPacket)) *UnknownPacketHandlerhandlePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnknownPacketHandlerhandlePacketCall) DoAndReturn(f func(receivedPacket)) *UnknownPacketHandlerhandlePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// setCloseError mocks base method.
func (m *MockUnknownPacketHandler) setCloseError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setCloseError", arg0)
}

// setCloseError indicates an expected call of setCloseError.
func (mr *MockUnknownPacketHandlerMockRecorder) setCloseError(arg0 interface{}) *UnknownPacketHandlersetCloseErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setCloseError", reflect.TypeOf((*MockUnknownPacketHandler)(nil).setCloseError), arg0)
	return &UnknownPacketHandlersetCloseErrorCall{Call: call}
}

// UnknownPacketHandlersetCloseErrorCall wrap *gomock.Call
type UnknownPacketHandlersetCloseErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnknownPacketHandlersetCloseErrorCall) Return() *UnknownPacketHandlersetCloseErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnknownPacketHandlersetCloseErrorCall) Do(f func(error)) *UnknownPacketHandlersetCloseErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnknownPacketHandlersetCloseErrorCall) DoAndReturn(f func(error)) *UnknownPacketHandlersetCloseErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
