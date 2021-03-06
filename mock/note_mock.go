// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zchary-ma/grpc-server/proto (interfaces: NoteServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/zchary-ma/grpc-server/proto"
	grpc "google.golang.org/grpc"
)

// MockNoteServiceClient is a mock of NoteServiceClient interface.
type MockNoteServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNoteServiceClientMockRecorder
}

// MockNoteServiceClientMockRecorder is the mock recorder for MockNoteServiceClient.
type MockNoteServiceClientMockRecorder struct {
	mock *MockNoteServiceClient
}

// NewMockNoteServiceClient creates a new mock instance.
func NewMockNoteServiceClient(ctrl *gomock.Controller) *MockNoteServiceClient {
	mock := &MockNoteServiceClient{ctrl: ctrl}
	mock.recorder = &MockNoteServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteServiceClient) EXPECT() *MockNoteServiceClientMockRecorder {
	return m.recorder
}

// CreateNote mocks base method.
func (m *MockNoteServiceClient) CreateNote(arg0 context.Context, arg1 *proto.Note, arg2 ...grpc.CallOption) (*proto.Id, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNote", varargs...)
	ret0, _ := ret[0].(*proto.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNote indicates an expected call of CreateNote.
func (mr *MockNoteServiceClientMockRecorder) CreateNote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNote", reflect.TypeOf((*MockNoteServiceClient)(nil).CreateNote), varargs...)
}

// DeleteNote mocks base method.
func (m *MockNoteServiceClient) DeleteNote(arg0 context.Context, arg1 *proto.IdSet, arg2 ...grpc.CallOption) (*proto.IdSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNote", varargs...)
	ret0, _ := ret[0].(*proto.IdSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockNoteServiceClientMockRecorder) DeleteNote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockNoteServiceClient)(nil).DeleteNote), varargs...)
}

// GetNote mocks base method.
func (m *MockNoteServiceClient) GetNote(arg0 context.Context, arg1 *proto.IdSet, arg2 ...grpc.CallOption) (*proto.NoteList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNote", varargs...)
	ret0, _ := ret[0].(*proto.NoteList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockNoteServiceClientMockRecorder) GetNote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockNoteServiceClient)(nil).GetNote), varargs...)
}

// UpdateNote mocks base method.
func (m *MockNoteServiceClient) UpdateNote(arg0 context.Context, arg1 *proto.Note, arg2 ...grpc.CallOption) (*proto.Note, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateNote", varargs...)
	ret0, _ := ret[0].(*proto.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockNoteServiceClientMockRecorder) UpdateNote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockNoteServiceClient)(nil).UpdateNote), varargs...)
}
