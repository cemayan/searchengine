package service

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ReadSvcTestSuite struct {
	suite.Suite
	config     config.GeneralConfig
	svc        *ReadService
	client     *redis.Client
	mockClient redismock.ClientMock
}

func (suite *ReadSvcTestSuite) SetupSuite() {
	config.Init(constants.ReadApi, "../../configs/read/config.yaml")
	suite.config = config.GetConfig(constants.ReadApi)
	suite.svc = NewReadService(constants.ReadApi)
	db, mock := redismock.NewClientMock()
	suite.mockClient = mock
	suite.client = db

}

func TestReadSvcTestSuite(t *testing.T) {
	suite.Run(t, new(ReadSvcTestSuite))
}

func (suite *ReadSvcTestSuite) Test_Start() {

}
