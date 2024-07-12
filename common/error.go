package common

import (
	"encoding/json"
	"github.com/cemayan/searchengine/types"
)

func GetError(msg string) []byte {
	err := types.Error{Msg: msg}
	marshal, _ := json.Marshal(err)
	return marshal
}
