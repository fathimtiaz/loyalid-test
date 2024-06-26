// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "loyalid-test/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserByAuth0Id provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetUserByAuth0Id(_a0 context.Context, _a1 string) (domain.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
