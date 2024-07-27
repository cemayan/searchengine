package service

import (
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/messaging"
	"github.com/cemayan/searchengine/protos"
	"github.com/cemayan/searchengine/protos/backendreq"
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/cemayan/searchengine/trie"
	"github.com/cemayan/searchengine/types"
	"github.com/sirupsen/logrus"
	"strings"
)

type WriteService struct {
	ProjectName constants.Project
}

// TODO: add return type
// prepareData creates a new record for each trie node
// ConvertForIndexing(golang) => [g:[go,gol,gola,golan,golang],go:[go,gol,gola,golan,golang],...]
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

// AddRecordMetadataToDb creates a record for given scraped results
// This methods call by grpc method
func (ws *WriteService) AddRecordMetadataToDb(req *backendreq.BackendRequest) {
	value := map[string]interface{}{}
	value["items"] = req.Items

	db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.RecordMetadata, req.GetRecord(), value, nil)
}

// mergeTheMaps merges between prevMap and currentMap
// prevMap: {go:{go:0}} currentMap: {go:{go:0,gol:0}} => {go:{go:0,gol:0}}
func (ws *WriteService) mergeTheMaps(prevMap map[string]interface{}, currentMap map[string]int) map[string]interface{} {
	for k, _ := range currentMap {
		if _, ok := prevMap[k]; !ok {
			prevMap[k] = 0
		}
	}
	return prevMap
}

// addRecordsToDb add records to db
// each trie object adds with initial value
// it means related record  is never selected before
// last db object for golang  => {g:{g:0,go:0,gol:0,gola:0,golan:0,golang:0},go:{},...}
func (ws *WriteService) addRecordsToDb(records map[string]map[string]int) {

	for k := range records {

		var err error
		currentMap := records[k]

		foundedRecords, err := db.SelectedDb(ws.ProjectName, constants.Read).Get(constants.Record, k, nil)
		if err != nil {
			logrus.Error(err)
			return
		}

		_db := constants.Str2Db[config.GetConfig(ws.ProjectName).Db.SelectedDb.Read]

		prevMap := map[string]interface{}{}

		// Since redis and mongodb return object is different we need to separate
		if _db == constants.Redis {

			castedFoundedRecords := foundedRecords.([]interface{})

			if len(castedFoundedRecords) > 0 {
				prevMap = castedFoundedRecords[0].(map[string]interface{})
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.Record, k, ws.mergeTheMaps(prevMap, currentMap), nil)
			} else {
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.Record, k, currentMap, nil)
			}

		} else if _db == constants.MongoDb {

			if foundedRecords != nil {
				prevMap = foundedRecords.(map[string]interface{})
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.Record, k, ws.mergeTheMaps(prevMap, currentMap), nil)
			} else {
				err = db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.Record, k, currentMap, nil)
			}

		}

		if err != nil {
			logrus.Error(err)
		}
	}
}

// increaseValue increases the value
func (ws *WriteService) increaseValue(prevMap map[string]interface{}, rec types.SelectionRequest) error {

	_db := constants.Str2Db[config.GetConfig(ws.ProjectName).Db.SelectedDb.Read]

	if _, ok := prevMap[rec.SelectedKey]; ok {

		// Since redis and mongodb return object is different we need to separate
		if _db == constants.Redis {
			flt := prevMap[rec.SelectedKey].(float64)
			prevMap[rec.SelectedKey] = int(flt) + 1
		} else if _db == constants.MongoDb {
			i32 := prevMap[rec.SelectedKey].(int32)
			prevMap[rec.SelectedKey] = int(i32) + 1
		}

		return db.SelectedDb(ws.ProjectName, constants.Write).Set(constants.Record, rec.Query, prevMap, nil)
	} else {
		return errors.New("record not found")
	}
	return nil
}

// Selection selects query according to clicked value on app side
// ex: let's assume that you are typing "go" and results looks below:
// [go,gol,gola,golan,golang] than you click the "golang"
// That operation means increased the current value that you clicked
// on db: =  // last db object for golang  => {go:{go:0,gol:0,gola:0,golan:0,golang:1},...}
func (ws *WriteService) Selection(rec types.SelectionRequest) error {
	foundedRecords, err := db.SelectedDb(ws.ProjectName, constants.Read).Get(constants.Record, rec.Query, nil)
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

func (ws *WriteService) PublishToNats(data string) {
	err := messaging.MessagingServer.Publish(constants.NatsEventsStream, protos.GetEvent([]byte(data), pb.EventType_RECORD_CREATED))
	if err != nil {
		logrus.Errorln("messaging server publish err", err)
	}
}

func NewWriteService(projectName constants.Project) *WriteService {
	return &WriteService{ProjectName: projectName}
}
