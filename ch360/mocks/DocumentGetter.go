// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ch360 "github.com/waives/surf/ch360"
import context "context"
import mock "github.com/stretchr/testify/mock"

// DocumentGetter is an autogenerated mock type for the DocumentGetter type
type DocumentGetter struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *DocumentGetter) GetAll(ctx context.Context) (ch360.DocumentList, error) {
	ret := _m.Called(ctx)

	var r0 ch360.DocumentList
	if rf, ok := ret.Get(0).(func(context.Context) ch360.DocumentList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ch360.DocumentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
