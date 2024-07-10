package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// config is used to reach config  whenever what you want
var config GeneralConfig

type Serve struct {
	Port int `yaml:"port"`
}

// GeneralConfig represents parsed yaml values
type GeneralConfig struct {
	Version     int    `yaml:"version"`
	Environment string `yaml:"environment"`
	Serve       Serve  `yaml:"serve"`
}

var (
	ConfigName  = "config"
	ConfigPaths = "configs"
)

// Init unmarshalls the yaml
func Init() {
	yamlFile := readYaml(ConfigName, ConfigPaths)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

// GetConfig returns the config
func GetConfig() GeneralConfig {
	return config
}
