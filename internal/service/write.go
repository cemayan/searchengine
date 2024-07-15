package service

import (
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/trie"
	"github.com/cemayan/searchengine/types"
	"github.com/sirupsen/logrus"
	"strings"
)

type WriteService struct {
	ProjectName constants.Project
}

// TODO: add return type
func (ws *WriteService) prepareData(data string) {
	_trie := trie.New()
	arr := strings.Split(data, " ")

	spaceIdxArr := make([]int, 0)

	for i := range len(data) {
		if data[i] == ' ' {
			spaceIdxArr = append(spaceIdxArr, i)
		}
	}

	if len(arr) > 1 {
		for i := range len(arr) - 1 {

			_data := data

			if i != 0 {
				_data = data[spaceIdxArr[i]:len(data)]
				_data = strings.TrimSpace(_data)
			}

			ws.addRecordsToDb(_trie.ConvertForIndexing(_data))
		}
	} else if len(arr) == 1 {
		ws.addRecordsToDb(_trie.ConvertForIndexing(data))
	}

}

func (ws *WriteService) mergeTheMaps(prevMap map[string]interface{}, currentMap map[string]int) map[string]interface{} {
	for k, _ := range currentMap {
		if _, ok := prevMap[k]; !ok {
			prevMap[k] = 0
		}
	}
	return prevMap
}
func (ws *WriteService) addRecordsToDb(records map[string]map[string]int) {
	for k := range records {

		var err error
		currentMap := records[k]

		foundedRecords, err := db.SelectedDb(ws.ProjectName, constants.Read).Get(k, nil)
		if err != nil {
			logrus.Error(err)
			return
		}

		_db := constants.Str2Db[config.GetConfig(ws.ProjectName).Db.SelectedDb.Read]

		prevMap := map[string]interface{}{}

		if _db == constants.Redis {

			castedFoundedRecords := foundedRecords.([]interface{})

			if len(castedFoundedRecords) > 0 {
				prevMap = castedFoundedRecords[0].(map[string]interface{})
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(k, ws.mergeTheMaps(prevMap, currentMap), nil)
			} else {
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(k, currentMap, nil)
			}

		} else if _db == constants.MongoDb {

			if foundedRecords != nil {
				prevMap = foundedRecords.(map[string]interface{})
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(k, ws.mergeTheMaps(prevMap, currentMap), nil)
			} else {
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(k, currentMap, nil)
			}

		}

		if err != nil {
			logrus.Error(err)
		}
	}
}

func (ws *WriteService) increaseValue(prevMap map[string]interface{}, rec types.SelectionRequest) error {
	if _, ok := prevMap[rec.SelectedKey]; ok {
		prevCounter := prevMap[rec.SelectedKey].(float64)
		prevMap[rec.SelectedKey] = prevCounter + 1

		return db.SelectedDb(ws.ProjectName, constants.Write).Set(rec.Query, prevMap, nil)
	} else {
		return errors.New("record not found")
	}
	return nil
}

func (ws *WriteService) Selection(rec types.SelectionRequest) error {
	foundedRecords, err := db.SelectedDb(ws.ProjectName, constants.Read).Get(rec.Query, nil)
	if err != nil {
		return nil
	}

	_db := constants.Str2Db[config.GetConfig(ws.ProjectName).Db.SelectedDb.Read]

	if _db == constants.Redis {

		castedFoundedRecords := foundedRecords.([]interface{})

		if len(castedFoundedRecords) > 0 {
			prevMap := castedFoundedRecords[0].(map[string]interface{})
			err = ws.increaseValue(prevMap, rec)
		}
	} else if _db == constants.MongoDb {
		if foundedRecords != nil {
			prevMap := foundedRecords.(map[string]interface{})
			err = ws.increaseValue(prevMap, rec)
		}
	}

	return nil
}

func (ws *WriteService) Start(data string) {
	ws.prepareData(data)
}

func NewWriteService(projectName constants.Project) *WriteService {
	return &WriteService{ProjectName: projectName}
}
