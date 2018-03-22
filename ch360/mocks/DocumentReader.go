// Code generated by mockery v1.0.0
package mocks

import ch360 "github.com/CloudHub360/ch360.go/ch360"
import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"

// DocumentReader is an autogenerated mock type for the DocumentReader type
type DocumentReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: ctx, documentId
func (_m *DocumentReader) Read(ctx context.Context, documentId string) error {
	ret := _m.Called(ctx, documentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, documentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadResult provides a mock function with given fields: ctx, documentId, mode
func (_m *DocumentReader) ReadResult(ctx context.Context, documentId string, mode ch360.ReadMode) (io.ReadCloser, error) {
	ret := _m.Called(ctx, documentId, mode)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, ch360.ReadMode) io.ReadCloser); ok {
		r0 = rf(ctx, documentId, mode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ch360.ReadMode) error); ok {
		r1 = rf(ctx, documentId, mode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}