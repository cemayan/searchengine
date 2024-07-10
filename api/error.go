package api

import (
	"encoding/json"
	"github.com/cemayan/searchengine/types"
)

func getError(msg string) []byte {
	err := types.Error{Msg: msg}
	marshal, _ := json.Marshal(err)
	return marshal
}
