package cli

import (
	"fmt"
	"os"
)

// Shows usage of all available commands and exits
func ShowUsage() {
	fmt.Fprintln(os.Stderr, "Usage: pm <command> [<args>]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  list		List projects")
	fmt.Fprintln(os.Stderr, "  go		Go to project")
	fmt.Fprintln(os.Stderr, "  remove	Remove project")
	fmt.Fprintln(os.Stderr, "  add 		Add project")
	os.Exit(1)
}

// Shows usage of a command and exits
func ShowCmdUsage(usageMsg string) {
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: pm %s", usageMsg))
	os.Exit(1)
}
