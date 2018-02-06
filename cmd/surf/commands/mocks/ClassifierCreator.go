// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// ClassifierCreator is an autogenerated mock type for the ClassifierCreator type
type ClassifierCreator struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (_m *ClassifierCreator) Create(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}