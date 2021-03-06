// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/waives/surf/ch360 (interfaces: DocumentCreator)

package mocks

import (
	context "context"
	pegomock "github.com/petergtz/pegomock"
	ch360 "github.com/waives/surf/ch360"
	io "io"
	"reflect"
	"time"
)

type MockDocumentCreator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockDocumentCreator(options ...pegomock.Option) *MockDocumentCreator {
	mock := &MockDocumentCreator{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockDocumentCreator) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockDocumentCreator) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockDocumentCreator) Create(ctx context.Context, fileContents io.Reader) (ch360.Document, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDocumentCreator().")
	}
	params := []pegomock.Param{ctx, fileContents}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Create", params, []reflect.Type{reflect.TypeOf((*ch360.Document)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 ch360.Document
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(ch360.Document)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockDocumentCreator) VerifyWasCalledOnce() *VerifierMockDocumentCreator {
	return &VerifierMockDocumentCreator{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockDocumentCreator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockDocumentCreator {
	return &VerifierMockDocumentCreator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockDocumentCreator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockDocumentCreator {
	return &VerifierMockDocumentCreator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockDocumentCreator) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockDocumentCreator {
	return &VerifierMockDocumentCreator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockDocumentCreator struct {
	mock                   *MockDocumentCreator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockDocumentCreator) Create(ctx context.Context, fileContents io.Reader) *MockDocumentCreator_Create_OngoingVerification {
	params := []pegomock.Param{ctx, fileContents}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Create", params, verifier.timeout)
	return &MockDocumentCreator_Create_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDocumentCreator_Create_OngoingVerification struct {
	mock              *MockDocumentCreator
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDocumentCreator_Create_OngoingVerification) GetCapturedArguments() (context.Context, io.Reader) {
	ctx, fileContents := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], fileContents[len(fileContents)-1]
}

func (c *MockDocumentCreator_Create_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []io.Reader) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]io.Reader, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(io.Reader)
		}
	}
	return
}
