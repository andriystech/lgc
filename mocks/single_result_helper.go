// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SingleResultHelper is an autogenerated mock type for the SingleResultHelper type
type SingleResultHelper struct {
	mock.Mock
}

// Decode provides a mock function with given fields: v
func (_m *SingleResultHelper) Decode(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
