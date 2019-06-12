// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import results "github.com/CloudHub360/ch360.go/ch360/results"

// DocumentClassifier is an autogenerated mock type for the DocumentClassifier type
type DocumentClassifier struct {
	mock.Mock
}

// Classify provides a mock function with given fields: ctx, documentId, classifierName
func (_m *DocumentClassifier) Classify(ctx context.Context, documentId string, classifierName string) (*results.ClassificationResult, error) {
	ret := _m.Called(ctx, documentId, classifierName)

	var r0 *results.ClassificationResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *results.ClassificationResult); ok {
		r0 = rf(ctx, documentId, classifierName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*results.ClassificationResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, documentId, classifierName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
