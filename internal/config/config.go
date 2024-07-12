package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// config is used to reach config  whenever what you want
var config map[string]GeneralConfig

func init() {
	config = make(map[string]GeneralConfig)
}

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
	ReadApiConfig  = "read"
	WriteApiConfig = "write"
	ConfigPaths    = "configs"
)

// Init unmarshalls the yaml
func Init(projectName string) {
	var genConfig GeneralConfig
	yamlFile := readYaml(fmt.Sprintf("%s/config", projectName), ConfigPaths)
	err := yaml.Unmarshal(yamlFile, &genConfig)
	config[projectName] = genConfig
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

// GetConfig returns the config
func GetConfig(projectName string) GeneralConfig {
	return config[projectName]
}
