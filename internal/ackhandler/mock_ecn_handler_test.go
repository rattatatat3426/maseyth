// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rattatatat3426/maseyth/internal/ackhandler (interfaces: ECNHandler)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package ackhandler -destination mock_ecn_handler_test.go github.com/rattatatat3426/maseyth/internal/ackhandler ECNHandler
//
// Package ackhandler is a generated GoMock package.
package ackhandler

import (
	reflect "reflect"

	protocol "github.com/rattatatat3426/maseyth/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockECNHandler is a mock of ECNHandler interface.
type MockECNHandler struct {
	ctrl     *gomock.Controller
	recorder *MockECNHandlerMockRecorder
}

// MockECNHandlerMockRecorder is the mock recorder for MockECNHandler.
type MockECNHandlerMockRecorder struct {
	mock *MockECNHandler
}

// NewMockECNHandler creates a new mock instance.
func NewMockECNHandler(ctrl *gomock.Controller) *MockECNHandler {
	mock := &MockECNHandler{ctrl: ctrl}
	mock.recorder = &MockECNHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECNHandler) EXPECT() *MockECNHandlerMockRecorder {
	return m.recorder
}

// HandleNewlyAcked mocks base method.
func (m *MockECNHandler) HandleNewlyAcked(arg0 []*packet, arg1, arg2, arg3 int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNewlyAcked", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HandleNewlyAcked indicates an expected call of HandleNewlyAcked.
func (mr *MockECNHandlerMockRecorder) HandleNewlyAcked(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewlyAcked", reflect.TypeOf((*MockECNHandler)(nil).HandleNewlyAcked), arg0, arg1, arg2, arg3)
}

// LostPacket mocks base method.
func (m *MockECNHandler) LostPacket(arg0 protocol.PacketNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LostPacket", arg0)
}

// LostPacket indicates an expected call of LostPacket.
func (mr *MockECNHandlerMockRecorder) LostPacket(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LostPacket", reflect.TypeOf((*MockECNHandler)(nil).LostPacket), arg0)
}

// Mode mocks base method.
func (m *MockECNHandler) Mode() protocol.ECN {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(protocol.ECN)
	return ret0
}

// Mode indicates an expected call of Mode.
func (mr *MockECNHandlerMockRecorder) Mode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockECNHandler)(nil).Mode))
}

// SentPacket mocks base method.
func (m *MockECNHandler) SentPacket(arg0 protocol.PacketNumber, arg1 protocol.ECN) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentPacket", arg0, arg1)
}

// SentPacket indicates an expected call of SentPacket.
func (mr *MockECNHandlerMockRecorder) SentPacket(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentPacket", reflect.TypeOf((*MockECNHandler)(nil).SentPacket), arg0, arg1)
}
