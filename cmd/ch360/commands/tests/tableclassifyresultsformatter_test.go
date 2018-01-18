package tests

import (
	"bytes"
	"fmt"
	"github.com/CloudHub360/ch360.go/ch360/types"
	"github.com/CloudHub360/ch360.go/cmd/ch360/commands"
	"github.com/CloudHub360/ch360.go/cmd/ch360/commands/mocks"
	"github.com/CloudHub360/ch360.go/test/generators"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"runtime"
	"strings"
	"testing"
)

type TableResultsFormatterSuite struct {
	suite.Suite
	output             *bytes.Buffer
	sut                *commands.TableClassifyResultsFormatter
	filename           string
	result             *types.ClassificationResult
	mockWriterProvider *mocks.WriterProvider
}

func (suite *TableResultsFormatterSuite) SetupTest() {
	suite.output = &bytes.Buffer{}
	suite.mockWriterProvider = &mocks.WriterProvider{}
	suite.mockWriterProvider.On("Provide", mock.Anything).Return(suite.output, nil)
	suite.sut = commands.NewTableClassifyResultsFormatter(suite.mockWriterProvider)

	suite.filename = generators.String("filename")
	suite.result = &types.ClassificationResult{
		DocumentType: generators.String("documenttype"),
		IsConfident:  false,
	}
}

func TestTableResultsWriterRunner(t *testing.T) {
	suite.Run(t, new(TableResultsFormatterSuite))
}

func (suite *TableResultsFormatterSuite) TestStart_Writes_Table_Header() {
	suite.sut.Start()

	header := fmt.Sprintf(commands.ClassifyOutputFormat, "FILE", "DOCUMENT TYPE", "CONFIDENT")
	assert.Equal(suite.T(), header, suite.output.String())
}

func (suite *TableResultsFormatterSuite) TestWrites_ResultWithCorrectFormat() {
	expectedOutput := "document1.tif                        documenttype                     true\n"
	filename := "document1.tif"
	result := &types.ClassificationResult{
		DocumentType: "documenttype",
		IsConfident:  true,
	}

	err := suite.sut.WriteResult(filename, result)

	require.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, suite.output.String())
}

func (suite *TableResultsFormatterSuite) TestWrites_Filename() {
	err := suite.sut.WriteResult(suite.filename, suite.result)

	require.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(suite.output.String(), suite.filename))
}

func (suite *TableResultsFormatterSuite) TestWrites_DocumentType() {
	err := suite.sut.WriteResult(suite.filename, suite.result)

	require.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(suite.output.String(), suite.result.DocumentType))
}

func (suite *TableResultsFormatterSuite) TestWrites_False_For_Not_IsConfident() {
	suite.result.IsConfident = false

	err := suite.sut.WriteResult(suite.filename, suite.result)

	require.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(suite.output.String(), "false"))
}

func (suite *TableResultsFormatterSuite) TestWrites_Filename_Only_When_It_Has_Path() {
	var filename string
	if runtime.GOOS == "windows" { //So tests can run on both Windows & Linux
		filename = `C:\folder\document1.tif`
	} else {
		filename = `/var/something/document1.tif`
	}

	expectedFilename := `document1.tif`

	err := suite.sut.WriteResult(filename, suite.result)

	require.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFilename, suite.output.String()[:len(expectedFilename)])
}