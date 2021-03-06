package tests

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/waives/surf/ch360"
	"github.com/waives/surf/cmd/surf/commands"
	"github.com/waives/surf/cmd/surf/commands/mocks"
	"github.com/waives/surf/net"
	"github.com/waives/surf/test/generators"
	"testing"
)

type CreateExtractorSuite struct {
	suite.Suite
	output          *bytes.Buffer
	creator         *mocks.ExtractorCreator
	sut             *commands.CreateExtractorCmd
	modulesTemplate *ch360.ExtractorTemplate
	extractorName   string
	ctx             context.Context
}

const modulesTemplateJson = `{"modules":[{"id":"waives.name"},{"id":"waives.date"}]}`

func (suite *CreateExtractorSuite) SetupTest() {
	suite.output = &bytes.Buffer{}
	suite.creator = new(mocks.ExtractorCreator)

	suite.modulesTemplate = aModulesTemplate()
	suite.extractorName = generators.String("extractor-name")
	suite.sut = &commands.CreateExtractorCmd{
		Creator:       suite.creator,
		ExtractorName: suite.extractorName,
		Template:      suite.modulesTemplate,
	}

	suite.creator.On("CreateFromModules", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	suite.ctx = context.Background()
}

func aModulesTemplate() *ch360.ExtractorTemplate {
	modulesTemplate, _ := ch360.NewModulesTemplateFromJson(bytes.NewBufferString(modulesTemplateJson))
	return modulesTemplate
}

func TestCreateExtractorSuiteRunner(t *testing.T) {
	suite.Run(t, new(CreateExtractorSuite))
}

func (suite *CreateExtractorSuite) ClearExpectedCalls() {
	suite.creator.ExpectedCalls = nil
}

func (suite *CreateExtractorSuite) TestCreateExtractor_Execute_Calls_Client_With_Correct_Args() {
	suite.sut.Execute(context.Background())

	suite.creator.AssertCalled(suite.T(), "CreateFromModules", suite.ctx, suite.extractorName, *suite.modulesTemplate)
}

func (suite *CreateExtractorSuite) TestCreateExtractor_Execute_Returns_Error_If_The_Extractor_Cannot_Be_Created() {
	expectedErr := errors.New("Error message")
	suite.ClearExpectedCalls()
	suite.creator.
		On("CreateFromModules", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedErr)

	receivedErr := suite.sut.Execute(context.Background())

	assert.Equal(suite.T(), expectedErr, receivedErr)
}

func (suite *CreateExtractorSuite) TestCreateExtractor_Execute_Returns_Useful_Information_From_A_DetailedErrorResponse() {
	detailedErrorResponse := aDetailedErrorResponse()
	suite.ClearExpectedCalls()
	suite.creator.
		On("CreateFromModules", mock.Anything, mock.Anything, mock.Anything).
		Return(detailedErrorResponse)
	expectedErrMsg := `Extractor creation failed with the following error: Invalid Extractor Template

Module (not found):
  The module waives.supplier_identity2 does not exist.

Module waives.supplier_identity:
  Parameter "provider": No argument was specified (specified "")
`

	receivedErr := suite.sut.Execute(context.Background())

	assert.Error(suite.T(), receivedErr, expectedErrMsg)
}

func aDetailedErrorResponse() *net.DetailedErrorResponse {
	return &net.DetailedErrorResponse{
		Title: "Invalid Extractor Template",
		Errors: []map[string]interface{}{
			{
				"module_id":      "",
				"messages":       []string{"The module waives.supplier_identity2 does not exist."},
				"path":           "waives.supplier_identity2",
				"argument_name":  nil,
				"argument_value": "",
			},
			{
				"module_id":      "waives.supplier_identity",
				"messages":       []string{"No argument was specified"},
				"path":           "modules[0].arguments.provider",
				"argument_name":  "provider",
				"argument_value": "",
			},
		},
		Status:   422,
		Instance: "/account/jK16_1URgUGxSo6yyWjHag/invalid-extractor-template/supplier-id",
		Type:     "https://docs.waives.io/reference#invalid-extractor-template",
	}
}
