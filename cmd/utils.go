package cmd

import (
	"fmt"
	"os"
)

func Terminate(msg string) {
	// Show error message
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func ShowUsage() {
	// Show usage message and exit
	fmt.Fprintln(os.Stderr, "usage: pm <command> [<args>]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  go		Go to Project")
	fmt.Fprintln(os.Stderr, "  remove	Remove Project")
	fmt.Fprintln(os.Stderr, "  add 		Add Project")
	os.Exit(1)
}

func ShowCmdUsage(usageMsg string) {
	// Show usage message and exit
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: gg %s", usageMsg))
	os.Exit(1)
}
