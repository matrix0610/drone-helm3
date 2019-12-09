package helm

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) TestHelmCommandDecodeSuccess() {
	cmd := helmCommand("")
	err := cmd.Decode("upgrade")
	suite.Require().Nil(err)

	suite.EqualValues(cmd, "upgrade")
}

func (suite *ConfigTestSuite) TestHelmCommandDecodeFailure() {
	cmd := helmCommand("")
	err := cmd.Decode("execute order 66")
	suite.EqualError(err, "unknown command 'execute order 66'. If specified, command must be upgrade, delete, lint, or help")
}
