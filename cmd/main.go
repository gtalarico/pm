package cmd

import (
	"os"

	"github.com/pkg/errors"
)

func Run() {

	config := readConfig()
	var args []string = os.Args[1:]

	if len(args) == 0 {
		ShowUsage()
	}

	cmd, posArgs := parseCmd(args)
	if cmd.NumArgs != len(posArgs) {
		ShowCmdUsage(cmd.UsageMsg)
	}

	cmd.Run(posArgs, config)
}

func InvalidCommand(cmdName string) {
	errMsg := recover() // recover() retrives error passed by panic
	if errMsg != nil {
		ShowUsage()
	}
}

func GetCommand(cmdName string) (command Command, err error) {
	// defer InvalidCommand(cmdName)
	for _, cmd := range Commands {
		if cmd.Name == cmdName {
			command = cmd
			return
		}
	}
	err = errors.New("invalid command")
	return
}

func parseCmd(args []string) (cmd Command, posArgs []string) {
	cmdName := args[0]
	cmd, err := GetCommand(cmdName)
	if err != nil {
		ShowUsage()
	}
	posArgs = args[1:]
	return

}
