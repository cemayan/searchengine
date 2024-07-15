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

func (rs *ReadService) Start(data *string) (*types.SearchResponse, error) {

	foundedRecords, err := db.SelectedDb(rs.ProjectName, constants.Read).Get(*data, nil)
	if err != nil {
		return nil, nil
	}

	db := constants.Str2Db[config.GetConfig(rs.ProjectName).Db.SelectedDb.Read]

	var genericMap map[string]interface{}

	if db == constants.Redis {
		castedFoundedRecords := foundedRecords.([]interface{})

		if len(castedFoundedRecords) != 0 {
			genericMap = castedFoundedRecords[0].(map[string]interface{})
		}
	} else if db == constants.MongoDb {
		genericMap = foundedRecords.(map[string]interface{})
	}

	intMap := map[string]int{}
	for key, value := range genericMap {
		if key != "_id" {
			if db == constants.Redis {
				flt := value.(float64)
				intMap[key] = int(flt)
			} else if db == constants.MongoDb {
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
