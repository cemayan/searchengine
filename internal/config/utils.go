package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

// readYaml reads the given yaml
func readYaml(path string) []byte {

	file, err := os.ReadFile(path)

	if err != nil {
		logrus.Errorln("Error reading file:", err)
		return nil
	}

	return file
}
