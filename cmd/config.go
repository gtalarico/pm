package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const CFG_FILENAME = ".go-project"

type Config struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func readConfig() Config {
	configPath := userHomePath() + fmt.Sprintf("/%s", CFG_FILENAME)
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		Terminate(err.Error())
	}

	var config Config
	parsingError := json.Unmarshal(configBytes, &config)

	if parsingError != nil {
		parsingErrorMsg := fmt.Sprintf("Invalid config format: %s", parsingError)
		Terminate(parsingErrorMsg)
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
