// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PersonDeleter is an autogenerated mock type for the PersonDeleter type
type PersonDeleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *PersonDeleter) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPersonDeleter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersonDeleter creates a new instance of PersonDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonDeleter(t mockConstructorTestingTNewPersonDeleter) *PersonDeleter {
	mock := &PersonDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
