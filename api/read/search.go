package read

import (
	"encoding/json"
	"github.com/cemayan/searchengine/common"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/trie"
	"github.com/cemayan/searchengine/types"
	"net/http"
)

func (srv *Server) GetQuery(w http.ResponseWriter, r *http.Request, params GetQueryParams) {

	resp := types.SearchResponse{}

	if params.Q == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("query parameter is required"))
		return
	}

	foundedRecords, err := db.Db.Get(*params.Q, nil)
	if err != nil {
		return
	}

	if foundedRecords != "" {
		var convertedStr []interface{}
		err = json.Unmarshal([]byte(foundedRecords), &convertedStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(common.GetError("json unmarshall failed"))
			return
		}

		resultMap := convertedStr[0]
		mm := resultMap.(map[string]interface{})

		for k, _ := range mm {
			resp = append(resp, k)
		}
	}

	if ok := json.NewEncoder(w).Encode(resp); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("json encode failed"))
	}

}

func (srv *Server) GetTestQuery(w http.ResponseWriter, r *http.Request, params GetTestQueryParams) {
	resp := types.SearchResponse{}

	if params.Q == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("query parameter is required"))
		return
	}

	_trie := trie.New()
	_trie.Insert("tea")
	_trie.Insert("ted")
	_trie.Insert("ten")
	_trie.Insert("tent")

	resp = _trie.SearchByPrefix(*params.Q)

	if ok := json.NewEncoder(w).Encode(resp); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("json encode failed"))
	}
}
