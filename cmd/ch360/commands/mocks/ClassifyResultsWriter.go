// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import types "github.com/CloudHub360/ch360.go/ch360/types"

// ClassifyResultsWriter is an autogenerated mock type for the ClassifyResultsWriter type
type ClassifyResultsWriter struct {
	mock.Mock
}

// Finish provides a mock function with given fields:
func (_m *ClassifyResultsWriter) Finish() {
	_m.Called()
}

// Start provides a mock function with given fields:
func (_m *ClassifyResultsWriter) Start() {
	_m.Called()
}

// WriteResult provides a mock function with given fields: filename, result
func (_m *ClassifyResultsWriter) WriteResult(filename string, result *types.ClassificationResult) error {
	ret := _m.Called(filename, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *types.ClassificationResult) error); ok {
		r0 = rf(filename, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
