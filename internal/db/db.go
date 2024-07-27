package db

import (
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db/mongodb"
	internalredis "github.com/cemayan/searchengine/internal/db/redis"
	"github.com/sirupsen/logrus"
)

// Db represents of global Db map
// ex: db.Db["readapi"]["redis"] gives DB methods that's selected
var Db map[constants.Project]map[constants.Db]DB

func init() {
	dbMap := make(map[constants.Project]map[constants.Db]DB)
	dbMap[constants.ReadApi] = make(map[constants.Db]DB)
	dbMap[constants.WriteApi] = make(map[constants.Db]DB)
	dbMap[constants.Scheduler] = make(map[constants.Db]DB)
	Db = dbMap
}

// DB represents methods that will be implemented
// Every DB must be implemented those methods
type DB interface {
	GetAll() interface{}
	Get(dbName constants.DbName, key string, params *[]string) (interface{}, error)
	Set(dbName constants.DbName, key string, value interface{}, params *[]string) error
}

// SelectedDb returns DB according to given information
func SelectedDb(project constants.Project, dbType constants.DbType) DB {

	if dbType == constants.Read {
		selectedReadDb := config.GetConfig(project).Db.SelectedDb.Read
		return Db[project][constants.Str2Db[selectedReadDb]]
	} else if dbType == constants.Write {
		selectedWriteDb := config.GetConfig(project).Db.SelectedDb.Write
		return Db[project][constants.Str2Db[selectedWriteDb]]
	}

	return nil
}

// Init initializes the DB
// According to given projectName DB instance will be created and set to global DB variable
// for reach where ever what you want.
func Init(projectName constants.Project) {
	db := config.GetConfig(projectName).Db
	schedulerConf := config.GetConfig(projectName).Scheduler

	if schedulerConf.Enabled {
		Db[projectName][constants.Redis] = internalredis.New(projectName)
		Db[projectName][constants.MongoDb] = mongodb.New(projectName)
	} else {
		if db.Cache.Name == constants.Db2Str[constants.Redis] {
			Db[projectName][constants.Redis] = internalredis.New(projectName)
		}

		if db.Persistent.Name == constants.Db2Str[constants.MongoDb] {
			Db[projectName][constants.MongoDb] = mongodb.New(projectName)
		}

		if db.SelectedDb.Read != "" {
			logrus.Println(fmt.Sprintf("%s initialized for read operations", db.SelectedDb.Read))
		}

		if db.SelectedDb.Write != "" {
			logrus.Println(fmt.Sprintf("%s initialized for write operations", db.SelectedDb.Write))
		}
	}

	if len(Db) == 0 {
		panic("there is no cache/persistent db")
	}
}
