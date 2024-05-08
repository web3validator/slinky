// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// PriceAggregator is an autogenerated mock type for the PriceAggregator type
type PriceAggregator struct {
	mock.Mock
}

// AggregatePrices provides a mock function with given fields:
func (_m *PriceAggregator) AggregatePrices() {
	_m.Called()
}

// GetPrices provides a mock function with given fields:
func (_m *PriceAggregator) GetPrices() map[string]*big.Float {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPrices")
	}

	var r0 map[string]*big.Float
	if rf, ok := ret.Get(0).(func() map[string]*big.Float); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Float)
		}
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *PriceAggregator) Reset() {
	_m.Called()
}

// SetProviderPrices provides a mock function with given fields: provider, prices
func (_m *PriceAggregator) SetProviderPrices(provider string, prices map[string]*big.Float) {
	_m.Called(provider, prices)
}

// NewPriceAggregator creates a new instance of PriceAggregator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceAggregator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceAggregator {
	mock := &PriceAggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
