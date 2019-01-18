package cmd

import (
	"fmt"
	"path/filepath"
)

type Command struct {
	Name     string
	NumArgs  int
	UsageMsg string
	Run      func([]string, Config)
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
	projectName := args[0]
	projectPath := args[1]
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		Terminate(err.Error())
	}
	project := Project{
		Name: projectName,
		Path: absPath,
	}
	// TODO: Remove first - requires refactoring findProject
	// project := findProject(projectName, config)

	newProjects := append(config.Projects, project)
	config.Projects = newProjects
	writeConfig(config)
	printProjects(newProjects)
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
		Name:     "list",
		NumArgs:  0, // pm list
		UsageMsg: "list",
		Run:      cmdList,
	},
	Command{
		Name:     "add",
		NumArgs:  2, // pm add <name> <path>
		UsageMsg: "add <project-name> <path>",
		Run:      cmdAdd,
	},
	Command{
		Name:     "remove",
		NumArgs:  1, // pm remove <query>
		UsageMsg: "remove <project-name>",
		Run:      cmdRemove,
	},
	Command{
		Name:     "go",
		NumArgs:  1, // pm go <project-name>
		UsageMsg: "go <project-name>",
		Run:      cmdGo,
	},
}

// Method: Command.Static()
// func (cmd *Command) Static() {
// 	if cmd.Name == "go" {
// 		cmdGo()
// 	}
// }
