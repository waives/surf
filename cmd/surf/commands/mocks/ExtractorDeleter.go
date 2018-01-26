// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// ExtractorDeleter is an autogenerated mock type for the ExtractorDeleter type
type ExtractorDeleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: name
func (_m *ExtractorDeleter) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}