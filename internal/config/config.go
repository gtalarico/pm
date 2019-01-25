package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const CFG_FILENAME = ".pm.json"

func WriteConfig(config Config) (err error) {
	path := ConfigFilepath()
	configJson, _ := json.MarshalIndent(config, "", " ")
	writeErr := ioutil.WriteFile(path, configJson, 0644)
	err = errors.Wrap(writeErr, path)
	return
}

func CreateConfig(path string) (cfg Config, err error) {
	projects := []*Project{}
	cfg = Config{projects}
	err = WriteConfig(cfg)
	return
}

func ReadConfig() (cfg Config, err error) {
	path := ConfigFilepath()
	configBytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		// Try creating new file in case of read error (eg. file not found)
		cfg, err = CreateConfig(path)
		return
	}
	parsingError := json.Unmarshal(configBytes, &cfg)
	if parsingError != nil {
		err = errors.Wrap(parsingError, path)
		return
	}
	return
}

// Gets user home path using environment variable '$HOME'
func UserHomePath() (path string) {
	path = os.Getenv("HOME")
	if path == "" {
		err, _ := fmt.Printf("Could not get home directory")
		panic(err)
	}
	return
}

// Gets the full config filepath
func ConfigFilepath() string {
	return UserHomePath() + fmt.Sprintf("/%s", CFG_FILENAME)
}
