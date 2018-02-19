// Code generated by mockery v1.0.0
package mocks

import formatters "github.com/CloudHub360/ch360.go/output/formatters"
import io "io"
import mock "github.com/stretchr/testify/mock"

// ResultsFormatter is an autogenerated mock type for the ResultsFormatter type
type ResultsFormatter struct {
	mock.Mock
}

// Flush provides a mock function with given fields: writer
func (_m *ResultsFormatter) Flush(writer io.Writer) error {
	ret := _m.Called(writer)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(writer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Format provides a mock function with given fields:
func (_m *ResultsFormatter) Format() formatters.OutputFormat {
	ret := _m.Called()

	var r0 formatters.OutputFormat
	if rf, ok := ret.Get(0).(func() formatters.OutputFormat); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(formatters.OutputFormat)
	}

	return r0
}

// WriteResult provides a mock function with given fields: writer, filename, result, options
func (_m *ResultsFormatter) WriteResult(writer io.Writer, filename string, result interface{}, options formatters.FormatOption) error {
	ret := _m.Called(writer, filename, result, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer, string, interface{}, formatters.FormatOption) error); ok {
		r0 = rf(writer, filename, result, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}