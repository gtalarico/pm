package cli

import (
	"fmt"
	"os"
)

// Shows usage of all available commands and exits
func ShowUsage() {
	usage :=
		`Usage:
	pm list
	pm go <project-name>
	pm add <project-name> <project-path>
	pm remove <project-name>`
	fmt.Fprintln(os.Stderr, usage)
}

// Shows usage of a command and exits
func ShowCmdUsage(usageMsg string) {
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: pm %s", usageMsg))
	os.Exit(1)
}
