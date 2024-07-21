package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	RepoUrl = "github.com/cemayan/searchengine"
)

// readYaml reads the given yaml
func readYaml(name string, configFolder string) []byte {

	configPath, err := getModuleFolder()

	if err != nil {
		configPath = configFolder
	} else {
		configPath = configPath + "/" + configFolder
	}

	currentEnv := "dev"

	if env, found := os.LookupEnv("READAPI_ENV"); found {
		currentEnv = env
	}

	currentFile := fmt.Sprintf("%s/%s_%s.yaml", configPath, name, currentEnv)

	file, err := os.ReadFile(currentFile)

	if err != nil {
		return nil
	}

	return file
}

// getModuleFolder returns library location that is downloaded
func getModuleFolder() (string, error) {
	cmdOut, err := exec.Command("go", "list", "-m", "-f", "'{{.Dir}}'", RepoUrl).Output()
	return strings.Trim(strings.TrimSpace(string(cmdOut)), "''"), err
}
