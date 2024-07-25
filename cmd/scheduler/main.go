package main

import (
	"flag"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/service"
)

var (
	configPath string // get config yaml path as flag
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path of config yaml")
	flag.Parse()
	// config initializer
	config.Init(constants.Scheduler, configPath)
	// db initializer
	db.Init(constants.Scheduler)
}

func main() {

	svc := service.NewSchedulerService()
	svc.Mongo2Redis()

}
