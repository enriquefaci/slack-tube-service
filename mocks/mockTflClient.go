// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/thoeni/go-tfl (interfaces: Client)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	go_tfl "github.com/thoeni/go-tfl"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) GetTubeStatus() ([]go_tfl.Report, error) {
	ret := _m.ctrl.Call(_m, "GetTubeStatus")
	ret0, _ := ret[0].([]go_tfl.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetTubeStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTubeStatus")
}

func (_m *MockClient) SetBaseURL(_param0 string) {
	_m.ctrl.Call(_m, "SetBaseURL", _param0)
}

func (_mr *_MockClientRecorder) SetBaseURL(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBaseURL", arg0)
}