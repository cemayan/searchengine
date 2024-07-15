package service

import (
	"encoding/json"
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/types"
	"sort"
)

type ReadService struct {
	ProjectName constants.Project
	resultMap   map[string]float64
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

	if foundedRecords != "" {
		var convertedStr []interface{}
		err = json.Unmarshal([]byte(foundedRecords), &convertedStr)

		if err != nil {
			return nil, errors.New("json unmarshall failed")
		}

		genericMap := convertedStr[0].(map[string]interface{})

		intMap := map[string]float64{}
		for key, value := range genericMap {
			intMap[key] = value.(float64)
		}

		rs.resultMap = intMap
		rs.sort()

	}
	return rs.result, nil
}

func NewReadService(projectName constants.Project) *ReadService {
	return &ReadService{projectName, make(map[string]float64), nil}
}
