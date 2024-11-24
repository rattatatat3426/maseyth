// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rattatatat3426/maseyth/internal/ackhandler (interfaces: SentPacketTracker)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package ackhandler -destination mock_sent_packet_tracker_test.go github.com/rattatatat3426/maseyth/internal/ackhandler SentPacketTracker
//
// Package ackhandler is a generated GoMock package.
package ackhandler

import (
	reflect "reflect"

	protocol "github.com/rattatatat3426/maseyth/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSentPacketTracker is a mock of SentPacketTracker interface.
type MockSentPacketTracker struct {
	ctrl     *gomock.Controller
	recorder *MockSentPacketTrackerMockRecorder
}

// MockSentPacketTrackerMockRecorder is the mock recorder for MockSentPacketTracker.
type MockSentPacketTrackerMockRecorder struct {
	mock *MockSentPacketTracker
}

// NewMockSentPacketTracker creates a new mock instance.
func NewMockSentPacketTracker(ctrl *gomock.Controller) *MockSentPacketTracker {
	mock := &MockSentPacketTracker{ctrl: ctrl}
	mock.recorder = &MockSentPacketTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentPacketTracker) EXPECT() *MockSentPacketTrackerMockRecorder {
	return m.recorder
}

// GetLowestPacketNotConfirmedAcked mocks base method.
func (m *MockSentPacketTracker) GetLowestPacketNotConfirmedAcked() protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLowestPacketNotConfirmedAcked")
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// GetLowestPacketNotConfirmedAcked indicates an expected call of GetLowestPacketNotConfirmedAcked.
func (mr *MockSentPacketTrackerMockRecorder) GetLowestPacketNotConfirmedAcked() *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLowestPacketNotConfirmedAcked", reflect.TypeOf((*MockSentPacketTracker)(nil).GetLowestPacketNotConfirmedAcked))
	return &SentPacketTrackerGetLowestPacketNotConfirmedAckedCall{Call: call}
}

// SentPacketTrackerGetLowestPacketNotConfirmedAckedCall wrap *gomock.Call
type SentPacketTrackerGetLowestPacketNotConfirmedAckedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall) Return(arg0 protocol.PacketNumber) *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall) Do(f func() protocol.PacketNumber) *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall) DoAndReturn(f func() protocol.PacketNumber) *SentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReceivedPacket mocks base method.
func (m *MockSentPacketTracker) ReceivedPacket(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedPacket", arg0)
}

// ReceivedPacket indicates an expected call of ReceivedPacket.
func (mr *MockSentPacketTrackerMockRecorder) ReceivedPacket(arg0 any) *SentPacketTrackerReceivedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockSentPacketTracker)(nil).ReceivedPacket), arg0)
	return &SentPacketTrackerReceivedPacketCall{Call: call}
}

// SentPacketTrackerReceivedPacketCall wrap *gomock.Call
type SentPacketTrackerReceivedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SentPacketTrackerReceivedPacketCall) Return() *SentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SentPacketTrackerReceivedPacketCall) Do(f func(protocol.EncryptionLevel)) *SentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SentPacketTrackerReceivedPacketCall) DoAndReturn(f func(protocol.EncryptionLevel)) *SentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
