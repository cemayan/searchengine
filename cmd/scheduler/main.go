package main

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/service"
)

func init() {
	config.Init(constants.Scheduler)
	db.Init(constants.Scheduler)
}

func main() {

	svc := service.NewSchedulerService()
	svc.Mongo2Redis()

}
