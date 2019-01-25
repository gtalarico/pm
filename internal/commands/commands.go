package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gtalarico/pm/internal/cli"
	"github.com/gtalarico/pm/internal/config"
	"github.com/pkg/errors"
)

func GetCommand(cmdName string) (command Command, err error) {
	for _, cmd := range Commands() {
		if cmd.Name == cmdName {
			command = cmd
			return
		}
	}
	err = errors.New("invalid command")
	return
}

func CommandList(args []string, config config.Config) {
	cli.PrintProjects(config.Projects)
}

func CommandGo(args []string, cfg config.Config) {
	projectName := args[0]
	project, err := config.GetOneProject(projectName, cfg)
	if err != nil {
		cli.Abort(errors.Wrap(err, projectName))
	} else {
		cli.Shell(project.Path)
	}
}

func CommandAdd(args []string, cfg config.Config) {
	projectName := args[0]
	projectPath := args[1]
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		cli.Abort(errors.Wrap(err, "invalid path"))
	}
	newProject := &config.Project{
		Name: projectName,
		Path: absPath,
	}
	project := config.MatchProjects(projectName, cfg)

	if project == nil {
		cfg.Projects = append(cfg.Projects, newProject)
	} else {
		project.Path = newProject.Path
	}
	writeError := config.WriteConfig(cfg)
	if writeError != nil {
		cli.Abort(writeError)
	}
	cli.PrintProjects(cfg.Projects)
}

func CommandRemove(args []string, cfg config.Config) {
	var projectToKeep []*config.Project
	projectName := args[0]
	matchedProject, err := config.GetOneProject(projectName, cfg)
	if err != nil {
		cli.Abort(errors.Wrap(err, projectName))
	} else {
		for _, project := range cfg.Projects {
			if project.Name != matchedProject.Name {
				projectToKeep = append(projectToKeep, project)
			}
		}
		cfg.Projects = projectToKeep

		promptMsg := fmt.Sprintf("Delete '%s' [Y/n]? ", matchedProject.Name)
		confirm := cli.ConfirmPrompt(promptMsg, true)
		if !confirm {
			os.Exit(0)
		}

		writeError := config.WriteConfig(cfg)
		if writeError != nil {
			cli.Abort(writeError)
		}
		cli.PrintProjects(projectToKeep)
	}
}

func Commands() (commands []Command) {
	commands = []Command{
		{
			Name:     "list",
			NumArgs:  0, // pm list
			UsageMsg: "list",
			Run:      CommandList,
		},
		{
			Name:     "add",
			NumArgs:  2, // pm add <name> <path>
			UsageMsg: "add <project-name> <path>",
			Run:      CommandAdd,
		},
		{
			Name:     "remove",
			NumArgs:  1, // pm remove <query>
			UsageMsg: "remove <project-name>",
			Run:      CommandRemove,
		},
		{
			Name:     "go",
			NumArgs:  1, // pm go <project-name>
			UsageMsg: "go <project-name>",
			Run:      CommandGo,
		},
	}
	return
}
