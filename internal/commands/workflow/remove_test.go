package workflow_test

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/konstellation-io/krt/pkg/krt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/konstellation-io/kli/cmd/config"
	"github.com/konstellation-io/kli/internal/commands/workflow"
	configservice "github.com/konstellation-io/kli/internal/services/product_configuration"
	"github.com/konstellation-io/kli/mocks"
)

type RemoveWorkflowSuite struct {
	suite.Suite

	renderer      *mocks.MockRenderer
	logger        *mocks.MockLogger
	handler       *workflow.Handler
	productConfig *configservice.ProductConfigService
	productName   string
}

func TestRemoveWorkflowSuite(t *testing.T) {
	suite.Run(t, new(RemoveWorkflowSuite))
}

func (s *RemoveWorkflowSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.logger = mocks.NewMockLogger(ctrl)
	renderer := mocks.NewMockRenderer(ctrl)
	mocks.AddLoggerExpects(s.logger)

	viper.SetDefault(config.KaiProductConfigFolder, ".kai")

	s.productName = "my-product"

	s.renderer = renderer

	s.handler = workflow.NewHandler(
		s.logger,
		s.renderer,
	)
}

func (s *RemoveWorkflowSuite) TearDownSuite(_, _ string) {
	err := os.RemoveAll(".kai")
	s.Require().NoError(err)
}

func (s *RemoveWorkflowSuite) AfterTest(_, _ string) {
	if err := os.RemoveAll(".kai"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			s.T().Fatalf("error cleaning tmp path: %s", err)
		}
	}
}

func (s *RemoveWorkflowSuite) BeforeTest(_, _ string) {
	err := s.productConfig.WriteConfiguration(
		_getDefaultKaiConfig(),
		s.productName,
	)
	s.Require().NoError(err)
}

func (s *RemoveWorkflowSuite) TestRemoveWorkflow_ExpectOk() {
	// GIVEN
	server := "server1"

	s.renderer.EXPECT().RenderWorkflows([]krt.Workflow{})

	// WHEN
	err := s.handler.RemoveWorkflow(&workflow.RemoveWorkflowOpts{
		ServerName: server,
		ProductID:  s.productName,
		WorkflowID: _getDefaultWorkflow().Name,
	})

	// THEN
	s.Require().NoError(err)
}

func (s *RemoveWorkflowSuite) TestRemoveWorkflow_NoExistingProduct_ExpectError() {
	// GIVEN
	server := "server1"

	// WHEN
	err := s.handler.RemoveWorkflow(&workflow.RemoveWorkflowOpts{
		ServerName: server,
		ProductID:  "some-product",
		WorkflowID: _getDefaultWorkflow().Name,
	})

	// THEN
	s.Require().Error(err)
	s.Require().ErrorIs(err, configservice.ErrProductConfigNotFound)
}

func (s *RemoveWorkflowSuite) TestRemoveWorkflow_NoExistingWorkflow_ExpectError() {
	// GIVEN
	server := "server1"

	// WHEN
	err := s.handler.RemoveWorkflow(&workflow.RemoveWorkflowOpts{
		ServerName: server,
		ProductID:  s.productName,
		WorkflowID: "some-workflow",
	})

	// THEN
	s.Require().Error(err)
	s.Require().ErrorIs(err, configservice.ErrWorkflowNotFound)
}
