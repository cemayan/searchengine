package db

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	internalredis "github.com/cemayan/searchengine/internal/db/redis"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DbTestSuite struct {
	suite.Suite
	config     config.GeneralConfig
	client     *redis.Client
	mockClient redismock.ClientMock
}

func (suite *DbTestSuite) SetupSuite() {
	config.Init(constants.ReadApi, "../../configs/read/config.yaml")
	db, mock := redismock.NewClientMock()
	suite.mockClient = mock
	suite.client = db
	suite.config = config.GetConfig(constants.ReadApi)

}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

func (suite *DbTestSuite) Test_Init_Redis() {

	internalredis.Client = suite.client
	suite.mockClient.ExpectPing().RedisNil()
	Init(constants.ReadApi)
	suite.NotNil(Db[constants.ReadApi][constants.Redis])
}

func (suite *DbTestSuite) Test_Init_MongoDb() {
	internalredis.Client = suite.client
	Init(constants.ReadApi)
	suite.NotNil(Db[constants.ReadApi][constants.MongoDb])
}

func (suite *DbTestSuite) Test_SelectedDb() {
	db := SelectedDb(constants.ReadApi, constants.Read)
	suite.Equal(Db[constants.ReadApi][constants.Redis], db)
}
