package config

import (
	"fmt"
	"github.com/cemayan/searchengine/constants"
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
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Scheduler struct {
	Enabled bool `yaml:"enabled"`
}

type Cache struct {
	Enabled bool `yaml:"enabled"`
}

type ScraperServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Scraper struct {
	Server ScraperServer `yaml:"server"`
}

// GeneralConfig represents parsed yaml values
type GeneralConfig struct {
	Version     int       `yaml:"version"`
	Environment string    `yaml:"environment"`
	Serve       Serve     `yaml:"serve"`
	Db          Db        `yaml:"db"`
	Scheduler   Scheduler `yaml:"scheduler"`
	Cache       Cache     `yaml:"cache"`
	Scraper     Scraper   `yaml:"scraper"`
}

var (
	ConfigPaths = "configs"
)

// Init unmarshalls the yaml
func Init(projectName constants.Project) {
	var genConfig GeneralConfig
	yamlFile := readYaml(fmt.Sprintf("%s/config", constants.ProjectMap[projectName]), ConfigPaths)
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
