package cmd

import (
	"fmt"
	"os"
)

type Command struct {
	Name    string
	NumArgs int
	Run     func([]string, Config)
}

func cmdList(args []string, config Config) {
	for _, project := range config.Projects {
		fmt.Println(project.Name)
	}
}

func cmdGo(args []string, config Config) {
	projectName := args[0]
	for _, project := range config.Projects {
		if projectName == project.Name {

			Shell(project.Path)
			os.Exit(0)
		}
	}
	Terminate("Project not found")
}

func cmdAdd(args []string, config Config) {
	fmt.Println("Cmd: Add")
	fmt.Println(args)
}

func cmdRemove(args []string, config Config) {
	fmt.Println("Cmd: Remove")
	fmt.Println(args)
}

var Commands = [...]Command{
	Command{
		Name:    "list",
		NumArgs: 0,
		Run:     cmdList,
	},
	Command{
		Name:    "add",
		NumArgs: 1,
		Run:     cmdAdd,
	},
	Command{
		Name:    "remove",
		NumArgs: 1,
		Run:     cmdRemove,
	},
	Command{
		Name:    "go",
		NumArgs: 1,
		Run:     cmdGo,
	},
}

// Method: Command.Static()
// func (cmd *Command) Static() {
// 	if cmd.Name == "go" {
// 		cmdGo()
// 	}
// }
