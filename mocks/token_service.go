// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/andriystech/lgc/models"
	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: _a0, _a1
func (_m *TokenService) GenerateToken(_a0 context.Context, _a1 *models.User) (*models.Token, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Token
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByToken provides a mock function with given fields: _a0, _a1
func (_m *TokenService) GetUserByToken(_a0 context.Context, _a1 string) (*models.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
