package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Project struct {
	Path string `json:"path"`
}

type Config struct {
	Projects []Project `json:"projects"`
}

func terminate(msg string) {
	// Show error message
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func showUsage() {
	// Show usage message and exit
	fmt.Fprint(os.Stderr, "Usage: gp <project-name>")
	os.Exit(1)
}

func homePath() (path string) {
	path = os.Getenv("HOME")
	if path == "" {
		err, _ := fmt.Printf("Invalid Home Dir:[%s]\n", path)
		panic(err)
	}
	return
}

func readConfig() {
	configPath := homePath() + "/.go-project"
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		terminate(err.Error())
	}

	var config Config
	parsingError := json.Unmarshal(configBytes, &config)

	if parsingError != nil {
		parsingErrorMsg := fmt.Sprintf("Invalid config format: %s", parsingError)
		terminate(parsingErrorMsg)
	}

	// fmt.Println(config.Projects)
	fmt.Println(config.Projects[0].Path)
}

func Run() {
	fmt.Println("STarted")
	var args []string = os.Args[1:]
	if len(args) == 0 {
		showUsage()
	}

	readConfig()
}
