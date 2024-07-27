package config

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) SetupSuite() {

}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) Test_Init() {
	Init(constants.ReadApi, "../../configs/read/config.yaml")
	suite.NotNil(config)
	suite.NotNil(config[constants.ReadApi])
	suite.Equal(1, config[constants.ReadApi].Version)
}
