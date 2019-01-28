package config

import (
	"os"
	"strings"
	"testing"
)

func TestUserHomePath(t *testing.T) {
	homePath := UserHomePath()
	if homePath != os.Getenv("HOME") {
		t.Error("Could not get HOME path")
	}
}

func TestConfigFilePath(t *testing.T) {
	path := configFilepath()
	endsWithConfigFile := strings.Contains(path, ConfigFilename)
	if !endsWithConfigFile {
		t.Error("Config filepath does not end with config filename: {}", path)
	}
}

// func WriteConfig(config Config) (err error) {
// func CreateConfig(path string) (cfg Config, err error) {
// func ReadConfig() (cfg Config, err error) {
