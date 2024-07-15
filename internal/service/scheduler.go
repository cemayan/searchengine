package service

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/db"
)

type SchedulerService struct {
}

func (ss *SchedulerService) Mongo2Redis() {
	allNames := db.Db[constants.Scheduler][constants.MongoDb].GetAll().([]string)

	for _, name := range allNames {
		get, _ := db.Db[constants.Scheduler][constants.MongoDb].Get(name, nil)
		err := db.Db[constants.Scheduler][constants.Redis].Set(name, get, nil)
		if err != nil {
			return
		}
	}
}

func NewSchedulerService() *SchedulerService {

	return &SchedulerService{}
}
