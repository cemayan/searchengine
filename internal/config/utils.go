package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

const (
	RepoUrl = "github.com/cemayan/searchengine"
)

// readYaml reads the given yaml
func readYaml(name string, configFolder string) []byte {

	currentPath, err := getModuleFolder()
	if err != nil {
		logrus.Errorf("fatal error config file: %w", err)
	}

	currentPath = currentPath + "/" + configFolder
	currentFile := currentPath + "/" + name + ".yaml"

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
