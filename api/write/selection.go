package write

import (
	"encoding/json"
	"fmt"
	"github.com/cemayan/searchengine/common"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (srv *Server) PostSelection(w http.ResponseWriter, r *http.Request) {
	rec := types.SelectionRequest{}
	err := json.NewDecoder(r.Body).Decode(&rec)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	foundedRecords, err := db.Db.Get(rec.Query, nil)
	if err != nil {
		return
	}

	if foundedRecords != "" {
		var convertedStr []interface{}
		err = json.Unmarshal([]byte(foundedRecords), &convertedStr)

		resultMap := convertedStr[0]
		prevMap := resultMap.(map[string]interface{})

		if _, ok := prevMap[rec.Selection]; ok {
			prevCounter := prevMap[rec.Selection].(float64)
			prevMap[rec.Selection] = prevCounter + 1

			recordJsonStr, _ := json.Marshal(prevMap)
			err = db.Db.Set(rec.Query, recordJsonStr, nil)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(common.GetError("record not found"))
			return
		}

		msg := fmt.Sprintf("record %s selected", rec.Selection)
		apiResp := types.ApiResponse{Msg: &msg}

		if ok := json.NewEncoder(w).Encode(apiResp); ok != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(common.GetError("json encode failed"))
		}
	}

}
