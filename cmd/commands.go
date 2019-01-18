package cmd

import (
	"fmt"
)

type Command struct {
	Name string
	Args int
	Run  func(args []string) (result bool, err error)
}

func cmdList(args []string) (result bool, err error) {
	fmt.Println("Cmd: List")
	fmt.Println(args)
	return true, nil
}

func cmdAdd(args []string) (result bool, err error) {
	fmt.Println("Cmd: Add")
	fmt.Println(args)
	return true, nil
}

func cmdRemove(args []string) (result bool, err error) {
	fmt.Println("Cmd: Remove")
	fmt.Println(args)
	return true, nil
}

func cmdGo(args []string) (result bool, err error) {
	fmt.Println("Cmd: GO")
	fmt.Println(args)
	return true, nil
}

var Commands = [...]Command{
	Command{
		Name: "list",
		Args: 0,
		Run:  cmdList,
	},
	Command{
		Name: "add",
		Args: 1,
		Run:  cmdAdd,
	},
	Command{
		Name: "remove",
		Args: 1,
		Run:  cmdRemove,
	},
	Command{
		Name: "go",
		Args: 1,
		Run:  cmdGo,
	},
}

// Method: Command.Static()
// func (cmd *Command) Static() {
// 	if cmd.Name == "go" {
// 		cmdGo()
// 	}
// }

func InvalidCommand(cmdName string) {
	errorMsg := recover() // recover() retrives error passed by panic
	if errorMsg != nil {
		terminate(fmt.Sprintf("%s: %s", errorMsg, cmdName))
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
