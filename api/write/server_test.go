package write

import (
	"bytes"
	"encoding/json"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	internalredis "github.com/cemayan/searchengine/internal/db/redis"
	"github.com/cemayan/searchengine/types"
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
	config.Init(constants.WriteApi, "../../configs/write/config_test.yaml")
	suite.config = config.GetConfig(constants.WriteApi)
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

	suite.Equal(3, len(routes))
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

	request, _ := http.NewRequest("GET", "http://localhost:8088/v1/health", nil)
	resp, err := http.DefaultClient.Do(request)
	suite.Nil(err)
	suite.Equal(resp.StatusCode, 200)

	cli, _ := redismock.NewClientMock()

	internalredis.Client = cli
	db.Init(constants.WriteApi)

	recordRequest := types.RecordRequest{Data: "alan turing"}
	byt, _ := json.Marshal(recordRequest)

	r := bytes.NewReader(byt)

	recordReq, _ := http.NewRequest("POST", "http://localhost:8088/v1/record", r)
	resp2, err := http.DefaultClient.Do(recordReq)
	suite.Nil(err)
	suite.Equal(resp2.StatusCode, 200)

	selectionRequest := types.SelectionRequest{
		Query:       "al",
		SelectedKey: "alan turing",
	}
	byt2, _ := json.Marshal(selectionRequest)

	r2 := bytes.NewReader(byt2)

	selectionReq, _ := http.NewRequest("POST", "http://localhost:8088/v1/selection", r2)
	resp3, err := http.DefaultClient.Do(selectionReq)
	suite.Nil(err)
	suite.Equal(resp3.StatusCode, 200)

}
