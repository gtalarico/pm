package cmd

import (
	"fmt"
	"os"
)

func terminate(msg string) {
	// Show error message
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func showUsage() {
	// Show usage message and exit
	fmt.Fprint(os.Stderr, "Usage: gp <project-name>")
	os.Exit(1)
}
