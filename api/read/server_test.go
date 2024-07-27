package read

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	internalredis "github.com/cemayan/searchengine/internal/db/redis"
	"github.com/go-redis/redismock/v9"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type ServerTestSuite struct {
	suite.Suite
	config config.GeneralConfig
	server Server
}

func (suite *ServerTestSuite) SetupSuite() {
	config.Init(constants.ReadApi, "../../configs/read/config.yaml")
	suite.config = config.GetConfig(constants.ReadApi)
	suite.server = NewServer()

}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (suite *ServerTestSuite) Test_configure() {
	suite.server.Configure()
	router := suite.server.router
	suite.NotNil(router)
	suite.NotNil(router.router)
	suite.NotNil(router.negroni)

	routes := map[string]struct{}{}

	router.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()

		if tpl != "/v1" {
			routes[tpl] = struct{}{}
		}

		return nil
	})

	suite.Equal(4, len(routes))
}

func (suite *ServerTestSuite) Test_corsHandler() {
	handler := corsHandler()
	request, _ := http.NewRequest("GET", "http://", nil)
	suite.Equal(true, handler.OriginAllowed(request))
}

func (suite *ServerTestSuite) Test_CheckEndpoints() {
	go func() {
		suite.server.Configure()
		suite.server.ListenAndServe()
	}()

	time.Sleep(3 * time.Second)

	request, _ := http.NewRequest("GET", "http://localhost:8087/v1/health", nil)
	resp, err := http.DefaultClient.Do(request)
	suite.Nil(err)
	suite.Equal(resp.StatusCode, 200)

	randReq, _ := http.NewRequest("GET", "http://localhost:8087/v1/random", nil)
	resp2, err := http.DefaultClient.Do(randReq)
	suite.Nil(err)
	suite.Equal(resp2.StatusCode, 404)

	cli, _ := redismock.NewClientMock()

	internalredis.Client = cli
	db.Init(constants.ReadApi)

	queryReq, _ := http.NewRequest("GET", "http://localhost:8087/v1/query?q=cem", nil)
	resp3, err := http.DefaultClient.Do(queryReq)
	suite.Nil(err)
	suite.Equal(resp3.StatusCode, 200)

	resultsReq, _ := http.NewRequest("GET", "http://localhost:8087/v1/results", nil)
	resultsReq.Header.Set(constants.XSearchEngineQuery, "test")
	resp4, err := http.DefaultClient.Do(resultsReq)
	suite.Nil(err)
	suite.Equal(resp4.StatusCode, 200)

}
