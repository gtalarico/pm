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
	fmt.Fprint(os.Stderr, "Usage: gg <project-name>")
	os.Exit(1)
}

func ShowCmdUsage(usageMsg string) {
	// Show usage message and exit
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: gg %s", usageMsg))
	os.Exit(1)
}
