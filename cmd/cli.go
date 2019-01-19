package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Fprintln(os.Stderr, "  list		List projects")
	fmt.Fprintln(os.Stderr, "  go		Go to project")
	fmt.Fprintln(os.Stderr, "  remove	Remove project")
	fmt.Fprintln(os.Stderr, "  add 		Add project")
	os.Exit(1)
}

func ShowCmdUsage(usageMsg string) {
	// Show usage message and exit
	fmt.Fprint(os.Stderr, fmt.Sprintf("Usage: pm %s", usageMsg))
	os.Exit(1)
}

func confirmPrompt(promptMsg string, default_ bool) bool {
	fmt.Print(promptMsg)

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	if response == "" {
		return default_
	}

	r := (strings.ToLower(response) == "y")
	return r
}
