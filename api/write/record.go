package write

import (
	"encoding/json"
	"fmt"
	"github.com/cemayan/searchengine/common"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/trie"
	"github.com/cemayan/searchengine/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (srv *Server) PostRecord(w http.ResponseWriter, r *http.Request) {

	rec := types.RecordRequest{}
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	_trie := trie.New()
	convertedData := _trie.ConvertForIndexing(rec.Data)

	for k := range convertedData {

		var err error

		val := convertedData[k]

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

	msg := fmt.Sprintf("%s record added to database successfully", rec.Data)
	apiResp := types.ApiResponse{Msg: &msg}

	w.Header().Set("Content-Type", "application/json")

	if ok := json.NewEncoder(w).Encode(apiResp); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("json encode failed"))
	}

}
