package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const CFG_FILENAME = ".go-project"

type Config struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Path string `json:"path"`
	Name string `json:"name"`
	// Command string `json:"command"`
}

func configFilepath() string {
	return userHomePath() + fmt.Sprintf("/%s", CFG_FILENAME)
}

func writeConfig(config Config) {
	path := configFilepath()
	configJson, _ := json.MarshalIndent(config, "", " ")
	writeErr := ioutil.WriteFile(path, configJson, 0644)
	if writeErr != nil {
		// writeErrMsg := fmt.Sprintf("Could not open config: %s", path)
		Terminate(errors.Wrap(writeErr, path))
	}
}

func readConfig() Config {
	path := configFilepath()
	configBytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		Terminate(readErr)
	}

	var config Config
	parsingError := json.Unmarshal(configBytes, &config)

	if parsingError != nil {
		Terminate(errors.Wrap(parsingError, path))
	}
	return config
}

func userHomePath() (path string) {
	path = os.Getenv("HOME")
	if path == "" {
		err, _ := fmt.Printf("Could not get home directory")
		panic(err)
	}
	return
}
