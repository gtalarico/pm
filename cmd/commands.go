package cmd

import (
	"fmt"
)

type Command struct {
	Name    string
	NumArgs int
	Run     func([]string, Config)
}

func printProjects(projects []Project) {
	for _, project := range projects {
		fmt.Println(project.Name)
	}
}

func cmdList(args []string, config Config) {
	printProjects(config.Projects)
}

func cmdGo(args []string, config Config) {
	projectName := args[0]
	project := findProject(projectName, config)
	Shell(project.Path)
}

func cmdAdd(args []string, config Config) {
	fmt.Println("Cmd: Add")
	fmt.Println(args)
}

func cmdRemove(args []string, config Config) {
	var projectToKeep []Project
	projectName := args[0]
	matchedProject := findProject(projectName, config)
	for _, project := range config.Projects {
		if project.Name != matchedProject.Name {
			projectToKeep = append(projectToKeep, project)
		}
	}
	config.Projects = projectToKeep
	writeConfig(config)
	printProjects(projectToKeep)
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
