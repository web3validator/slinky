// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WebSocketConnHandler is an autogenerated mock type for the WebSocketConnHandler type
type WebSocketConnHandler struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *WebSocketConnHandler) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dial provides a mock function with given fields: url
func (_m *WebSocketConnHandler) Dial(url string) error {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Dial")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields:
func (_m *WebSocketConnHandler) Read() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: message
func (_m *WebSocketConnHandler) Write(message []byte) error {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWebSocketConnHandler creates a new instance of WebSocketConnHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebSocketConnHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebSocketConnHandler {
	mock := &WebSocketConnHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
