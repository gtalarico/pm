package commands

import "errors"

func GetCommand(cmdName string) (command Command, err error) {
	for _, cmd := range Commands {
		if cmd.Name == cmdName {
			command = cmd
			return
		}
	}
	err = errors.New("invalid command")
	return
}
