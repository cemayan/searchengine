package service

import (
	"encoding/json"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/trie"
	"github.com/sirupsen/logrus"
	"strings"
)

type WriteService struct{}

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

func (ws *WriteService) addRecordsToDb(records map[string]map[string]int) {
	for k := range records {

		var err error
		val := records[k]

		foundedRecords, err := db.Db.Get(k, nil)
		if err != nil {
			return
		}

		if foundedRecords != "" {
			var convertedStr []interface{}
			err = json.Unmarshal([]byte(foundedRecords), &convertedStr)

			resultMap := convertedStr[0]
			prevMap := resultMap.(map[string]interface{})

			for k, _ := range val {
				if _, ok := prevMap[k]; !ok {
					prevMap[k] = 0
				}
			}

			recordJsonStr, _ := json.Marshal(prevMap)
			err = db.Db.Set(k, recordJsonStr, nil)
		} else {
			recordJsonStr, _ := json.Marshal(val)
			err = db.Db.Set(k, recordJsonStr, nil)
		}

		if err != nil {
			logrus.Error(err)
		}
	}
}

func (ws *WriteService) Start(data string) {
	ws.prepareData(data)
}
