// Code generated by MockGen. DO NOT EDIT.
// Source: net (interfaces: PacketConn)
//
// Generated by this command:
//
//	mockgen -typed -package quic -self_package github.com/rattatatat3426/maseyth -self_package github.com/rattatatat3426/maseyth -destination mock_packetconn_test.go net PacketConn
//

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockPacketConn is a mock of PacketConn interface.
type MockPacketConn struct {
	ctrl     *gomock.Controller
	recorder *MockPacketConnMockRecorder
}

// MockPacketConnMockRecorder is the mock recorder for MockPacketConn.
type MockPacketConnMockRecorder struct {
	mock *MockPacketConn
}

// NewMockPacketConn creates a new mock instance.
func NewMockPacketConn(ctrl *gomock.Controller) *MockPacketConn {
	mock := &MockPacketConn{ctrl: ctrl}
	mock.recorder = &MockPacketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketConn) EXPECT() *MockPacketConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPacketConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPacketConnMockRecorder) Close() *MockPacketConnCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketConn)(nil).Close))
	return &MockPacketConnCloseCall{Call: call}
}

// MockPacketConnCloseCall wrap *gomock.Call
type MockPacketConnCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnCloseCall) Return(arg0 error) *MockPacketConnCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnCloseCall) Do(f func() error) *MockPacketConnCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnCloseCall) DoAndReturn(f func() error) *MockPacketConnCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockPacketConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockPacketConnMockRecorder) LocalAddr() *MockPacketConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockPacketConn)(nil).LocalAddr))
	return &MockPacketConnLocalAddrCall{Call: call}
}

// MockPacketConnLocalAddrCall wrap *gomock.Call
type MockPacketConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnLocalAddrCall) Return(arg0 net.Addr) *MockPacketConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnLocalAddrCall) Do(f func() net.Addr) *MockPacketConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnLocalAddrCall) DoAndReturn(f func() net.Addr) *MockPacketConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadFrom mocks base method.
func (m *MockPacketConn) ReadFrom(arg0 []byte) (int, net.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(net.Addr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFrom indicates an expected call of ReadFrom.
func (mr *MockPacketConnMockRecorder) ReadFrom(arg0 any) *MockPacketConnReadFromCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockPacketConn)(nil).ReadFrom), arg0)
	return &MockPacketConnReadFromCall{Call: call}
}

// MockPacketConnReadFromCall wrap *gomock.Call
type MockPacketConnReadFromCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnReadFromCall) Return(arg0 int, arg1 net.Addr, arg2 error) *MockPacketConnReadFromCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnReadFromCall) Do(f func([]byte) (int, net.Addr, error)) *MockPacketConnReadFromCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnReadFromCall) DoAndReturn(f func([]byte) (int, net.Addr, error)) *MockPacketConnReadFromCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDeadline mocks base method.
func (m *MockPacketConn) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockPacketConnMockRecorder) SetDeadline(arg0 any) *MockPacketConnSetDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetDeadline), arg0)
	return &MockPacketConnSetDeadlineCall{Call: call}
}

// MockPacketConnSetDeadlineCall wrap *gomock.Call
type MockPacketConnSetDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnSetDeadlineCall) Return(arg0 error) *MockPacketConnSetDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnSetDeadlineCall) Do(f func(time.Time) error) *MockPacketConnSetDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnSetDeadlineCall) DoAndReturn(f func(time.Time) error) *MockPacketConnSetDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockPacketConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockPacketConnMockRecorder) SetReadDeadline(arg0 any) *MockPacketConnSetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetReadDeadline), arg0)
	return &MockPacketConnSetReadDeadlineCall{Call: call}
}

// MockPacketConnSetReadDeadlineCall wrap *gomock.Call
type MockPacketConnSetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnSetReadDeadlineCall) Return(arg0 error) *MockPacketConnSetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnSetReadDeadlineCall) Do(f func(time.Time) error) *MockPacketConnSetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnSetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *MockPacketConnSetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockPacketConn) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockPacketConnMockRecorder) SetWriteDeadline(arg0 any) *MockPacketConnSetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetWriteDeadline), arg0)
	return &MockPacketConnSetWriteDeadlineCall{Call: call}
}

// MockPacketConnSetWriteDeadlineCall wrap *gomock.Call
type MockPacketConnSetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnSetWriteDeadlineCall) Return(arg0 error) *MockPacketConnSetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnSetWriteDeadlineCall) Do(f func(time.Time) error) *MockPacketConnSetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnSetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *MockPacketConnSetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WriteTo mocks base method.
func (m *MockPacketConn) WriteTo(arg0 []byte, arg1 net.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockPacketConnMockRecorder) WriteTo(arg0, arg1 any) *MockPacketConnWriteToCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockPacketConn)(nil).WriteTo), arg0, arg1)
	return &MockPacketConnWriteToCall{Call: call}
}

// MockPacketConnWriteToCall wrap *gomock.Call
type MockPacketConnWriteToCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketConnWriteToCall) Return(arg0 int, arg1 error) *MockPacketConnWriteToCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketConnWriteToCall) Do(f func([]byte, net.Addr) (int, error)) *MockPacketConnWriteToCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketConnWriteToCall) DoAndReturn(f func([]byte, net.Addr) (int, error)) *MockPacketConnWriteToCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
