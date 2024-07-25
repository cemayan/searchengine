package service

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/types"
	"sort"
)

type ReadService struct {
	ProjectName constants.Project
	resultMap   map[string]int
	result      *types.SearchResponse
}

// sort sorts the resultMap
// ex: {alan turi:0,alan turing: 5 }
// after sorting  {alan turing:5,alan turi: 0 }
func (rs *ReadService) sort() {
	keys := make([]string, 0, len(rs.resultMap))

	for k := range rs.resultMap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return rs.resultMap[keys[i]] > rs.resultMap[keys[j]]
	})

	rs.result = &keys
}

// GetResults returns an array  according to given query
// ex: alan turing
// {list:[title:"",url:""],...}
func (rs *ReadService) GetResults(query string) map[string]interface{} {
	currentDb := constants.Str2Db[config.GetConfig(rs.ProjectName).Db.SelectedDb.Read]
	foundedRecords, err := db.SelectedDb(rs.ProjectName, constants.Read).Get(constants.RecordMetadata, query, nil)
	if err != nil || foundedRecords == nil {
		return nil
	}

	var genericMap map[string]interface{}

	// Since redis and mongodb return object is different we need to separate
	if currentDb == constants.Redis {
		castedFoundedRecords := foundedRecords.([]interface{})

		if len(castedFoundedRecords) != 0 {
			genericMap = castedFoundedRecords[0].(map[string]interface{})
		}
	} else if currentDb == constants.MongoDb {
		genericMap = foundedRecords.(map[string]interface{})
	}

	return genericMap
}

// Start executes read operations for given data on selected DB
func (rs *ReadService) Start(data *string) (*types.SearchResponse, error) {

	currentDb := constants.Str2Db[config.GetConfig(rs.ProjectName).Db.SelectedDb.Read]

	foundedRecords, err := db.SelectedDb(rs.ProjectName, constants.Read).Get(constants.Record, *data, nil)
	if err != nil || foundedRecords == nil {
		return nil, err
	}

	var genericMap map[string]interface{}

	if currentDb == constants.Redis {
		castedFoundedRecords := foundedRecords.([]interface{})

		if len(castedFoundedRecords) != 0 {
			genericMap = castedFoundedRecords[0].(map[string]interface{})
		}
	} else if currentDb == constants.MongoDb {
		genericMap = foundedRecords.(map[string]interface{})
	}

	intMap := map[string]int{}
	for key, value := range genericMap {
		if key != "_id" {
			if currentDb == constants.Redis {
				flt := value.(float64)
				intMap[key] = int(flt)
			} else if currentDb == constants.MongoDb {
				i32 := value.(int32)
				intMap[key] = int(i32)
			}
		}
	}

	rs.resultMap = intMap
	rs.sort()

	return rs.result, nil
}

func NewReadService(projectName constants.Project) *ReadService {
	return &ReadService{projectName, make(map[string]int), nil}
}
