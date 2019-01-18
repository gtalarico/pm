package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
	// fmt.Println(config.Projects[0].Path)
}

func Run() {
	var args []string = os.Args[1:]
	if len(args) == 0 {
		showUsage()
	}
	cmd, posArgs := parseCmd(args)
	if cmd.Args != len(posArgs) {
		showUsage()
	}
	cmd.Run(posArgs)
	// readConfig()
}
