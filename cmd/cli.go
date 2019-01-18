package cmd

import (
	"fmt"
	"os"
)

func Terminate(err error) {
	// Show error message
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
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
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: pm %s", usageMsg))
	os.Exit(1)
}
