package cmd

import (
	"fmt"
	"os"
)

func Run() {

	config := readConfig()
	var args []string = os.Args[1:]

	if len(args) == 0 {
		ShowUsage()
	}

	cmd, posArgs := parseCmd(args)
	if cmd.NumArgs != len(posArgs) {
		ShowUsage()
	}

	cmd.Run(posArgs, config)
}

func InvalidCommand(cmdName string) {
	errMsg := recover() // recover() retrives error passed by panic
	if errMsg != nil {
		Terminate(fmt.Sprintf("%s: %s", errMsg, cmdName))
	}
}

func GetCommand(cmdName string) Command {
	defer InvalidCommand(cmdName)
	for _, cmd := range Commands {
		if cmd.Name == cmdName {
			return cmd
		}
	}
	panic("Invalid Command Error")
}

func parseCmd(args []string) (cmd Command, posArgs []string) {
	cmdName := args[0]
	cmd = GetCommand(cmdName)
	posArgs = args[1:]
	return

}
