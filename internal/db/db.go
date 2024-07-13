package db

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db/redis"
)

var Db DB

type DB interface {
	Get(key string, params []string) (string, error)
	Set(key string, value interface{}, params *[]string) error
}

func Init(projectName string) {
	db := config.GetConfig(projectName).Db
	if db.Name == constants.Redis {
		Db = redis.New(projectName)
	}
}
