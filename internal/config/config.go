package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

const ConfigFilename = ".pm.json"

func WriteConfig(config Config) (err error) {
	path := configFilepath()
	configJSON, _ := json.MarshalIndent(config, "", " ")
	writeErr := ioutil.WriteFile(path, configJSON, 0644)
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
	path := configFilepath()
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

// Gets user home path using environment variable '$HOME' or '%HOMEPATH%'
func UserHomePath() (path string) {
	if runtime.GOOS == "windows" {
		path = os.Getenv("HOMEPATH")
	} else {
		path = os.Getenv("HOME")
	}

	if path == "" {
		err, _ := fmt.Printf("Could not get home directory")
		panic(err)
	}
	return
}

// Gets the full config filepath
func configFilepath() string {
	return UserHomePath() + fmt.Sprintf("/%s", ConfigFilename)
}
