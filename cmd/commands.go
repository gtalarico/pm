package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
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
	project, err := getOneProject(projectName, config, false)
	if err != nil {
		Terminate(errors.Wrap(err, projectName))
	} else {
		Shell(project.Path)
	}
}

func cmdAdd(args []string, config Config) {
	projectName := args[0]
	projectPath := args[1]
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		Terminate(errors.Wrap(err, "invalid path"))
	}
	newProject := Project{
		Name: projectName,
		Path: absPath,
	}
	projects := searchProjects(projectName, config, true)
	if len(projects) == 0 {
		config.Projects = append(config.Projects, newProject)
	} else {
		for i, project := range config.Projects {
			if project.Name == newProject.Name {
				config.Projects[i] = newProject
			}
		}
	}
	writeConfig(config)
	printProjects(config.Projects)
}

func cmdRemove(args []string, config Config) {
	var projectToKeep []Project
	projectName := args[0]
	matchedProject, err := getOneProject(projectName, config, false)
	if err != nil {
		Terminate(errors.Wrap(err, projectName))
	} else {
		for _, project := range config.Projects {
			if project.Name != matchedProject.Name {
				projectToKeep = append(projectToKeep, project)
			}
		}
		config.Projects = projectToKeep

		promptMsg := fmt.Sprintf("Delete '%s' [Y/n]? ", matchedProject.Name)
		confirm := confirmPrompt(promptMsg, true)
		if confirm == false {
			Terminate(nil)
		}

		writeConfig(config)
		printProjects(projectToKeep)
	}
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
