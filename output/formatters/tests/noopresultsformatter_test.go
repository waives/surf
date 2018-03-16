package tests

import (
	"bytes"
	"github.com/CloudHub360/ch360.go/output/formatters"
	"github.com/CloudHub360/ch360.go/test/generators"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_Formatter_Copies_From_Reader_To_Writer(t *testing.T) {
	// Arrange
	srcBuffer := generators.Bytes()
	readerCloser := ioutil.NopCloser(bytes.NewBuffer(srcBuffer))
	targetBuffer := &bytes.Buffer{}
	sut := &formatters.NoopResultsFormatter{}

	// Act
	sut.WriteResult(targetBuffer, "", readerCloser, 0)

	// Assert
	assert.Equal(t, srcBuffer, targetBuffer.Bytes())
}

func Test_Formatter_Checks_Type_Of_Input_Result(t *testing.T) {
	// Arrange
	targetBuffer := &bytes.Buffer{}
	sut := &formatters.NoopResultsFormatter{}
	result := generators.Bytes() // any old type

	// Act
	err := sut.WriteResult(targetBuffer, "", result, 0)

	// Assert
	assert.EqualError(t, formatters.ErrUnexpectedType(result), err.Error())
}
