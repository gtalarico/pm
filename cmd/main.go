package cmd

import (
	"os"

	"github.com/gtalarico/pm/internal/cli"
	"github.com/gtalarico/pm/internal/commands"
	"github.com/gtalarico/pm/internal/config"
)

func Run() {

	// Get all args, excluding binary name
	var args []string = os.Args[1:]
	cmdName, posArgs, err := cli.ValidateArgs(args)

	// No Args
	if err != nil {
		cli.ShowUsage()
	}

	// Invalid command name
	// or
	// Not the right number of args for a given command
	cmd, err := commands.GetCommand(cmdName)
	if cmd.NumArgs != len(posArgs) {
		cli.ShowCmdUsage(cmd.UsageMsg)
	}

	// Get Config
	config, err := config.ReadConfig()
	if err != nil {
		cli.Terminate(err)
	}

	// Run Command
	cmd.Run(posArgs, config)

}
