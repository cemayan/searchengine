package config

import (
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// config is used to reach config  whenever what you want
var config map[constants.Project]GeneralConfig

func init() {
	config = make(map[constants.Project]GeneralConfig)
}

type Serve struct {
	Port int `yaml:"port"`
}

type SelectedDb struct {
	Read  string `yaml:"read"`
	Write string `yaml:"write"`
}

type Db struct {
	Cache      DbConfig   `yaml:"cache"`
	Persistent DbConfig   `yaml:"persistent"`
	SelectedDb SelectedDb `yaml:"selectedDb"`
}

type DbConfig struct {
	Name string `yaml:"name"`
	Uri  string `yaml:"uri"`
	Rs   string `yaml:"rs"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Scheduler struct {
	Enabled bool `yaml:"enabled"`
}

type CacheDirect struct {
	Enabled bool `yaml:"enabled"`
}

type ScraperServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Scraper struct {
	Server ScraperServer `yaml:"server"`
}

type Web struct {
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

type Consumer struct {
	Name   string `yaml:"name"`
	Stream string `yaml:"stream"`
}

type Nats struct {
	Url         string     `yaml:"url"`
	Streams     []string   `yaml:"streams"`
	Consumers   []Consumer `yaml:"consumers"`
	IsJsEnabled bool       `yaml:"js"`
}

type Kafka struct {
}

type Messaging struct {
	Nats  *Nats  `yaml:"nats"`
	Kafka *Kafka `yaml:"kafka"`
}

// GeneralConfig represents parsed yaml values
type GeneralConfig struct {
	Version     int         `yaml:"version"`
	Environment string      `yaml:"environment"`
	Serve       Serve       `yaml:"serve"`
	Db          Db          `yaml:"db"`
	Scheduler   Scheduler   `yaml:"scheduler"`
	CacheDirect CacheDirect `yaml:"cacheDirect"`
	Scraper     Scraper     `yaml:"scraper"`
	Web         Web         `yaml:"web"`
	Messaging   Messaging   `yaml:"messaging"`
}

// Init unmarshalls the yaml
func Init(projectName constants.Project, yamlPath string) {
	var genConfig GeneralConfig
	logrus.Println("yamlPath:", yamlPath)
	yamlFile := readYaml(yamlPath)
	err := yaml.Unmarshal(yamlFile, &genConfig)
	config[projectName] = genConfig
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

// GetConfig returns the config
func GetConfig(projectName constants.Project) GeneralConfig {
	return config[projectName]
}
