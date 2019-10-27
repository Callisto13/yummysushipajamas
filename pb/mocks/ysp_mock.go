// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Callisto13/yummysushipajamas/pb (interfaces: Basic_PrimeServer)

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	context "context"
	pb "github.com/Callisto13/yummysushipajamas/pb"
	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockBasic_PrimeServer is a mock of Basic_PrimeServer interface
type MockBasic_PrimeServer struct {
	ctrl     *gomock.Controller
	recorder *MockBasic_PrimeServerMockRecorder
}

// MockBasic_PrimeServerMockRecorder is the mock recorder for MockBasic_PrimeServer
type MockBasic_PrimeServerMockRecorder struct {
	mock *MockBasic_PrimeServer
}

// NewMockBasic_PrimeServer creates a new mock instance
func NewMockBasic_PrimeServer(ctrl *gomock.Controller) *MockBasic_PrimeServer {
	mock := &MockBasic_PrimeServer{ctrl: ctrl}
	mock.recorder = &MockBasic_PrimeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBasic_PrimeServer) EXPECT() *MockBasic_PrimeServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBasic_PrimeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBasic_PrimeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBasic_PrimeServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBasic_PrimeServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBasic_PrimeServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBasic_PrimeServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBasic_PrimeServer) Send(arg0 *pb.Resp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBasic_PrimeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBasic_PrimeServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBasic_PrimeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBasic_PrimeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBasic_PrimeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBasic_PrimeServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBasic_PrimeServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBasic_PrimeServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBasic_PrimeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBasic_PrimeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBasic_PrimeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBasic_PrimeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBasic_PrimeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBasic_PrimeServer)(nil).SetTrailer), arg0)
}
