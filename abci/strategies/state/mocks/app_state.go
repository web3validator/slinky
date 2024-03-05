// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AppState is an autogenerated mock type for the AppState type
type AppState struct {
	mock.Mock
}

// VerifyVoteExtensionState provides a mock function with given fields: ctx
func (_m *AppState) VerifyVoteExtensionState(ctx types.Context) (types.Context, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for VerifyVoteExtensionState")
	}

	var r0 types.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (types.Context, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) types.Context); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Context)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAppState creates a new instance of AppState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppState(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppState {
	mock := &AppState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}