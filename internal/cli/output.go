package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gtalarico/pm/internal/config"
)

func Abort(err error) {
	// Show error message and exit with error
	fmt.Fprint(os.Stderr, err.Error())
	os.Exit(1)
}

// Prompts user to confirm a "y" or "n" and returns as boolean
func ConfirmPrompt(promptMsg string, default_ bool) bool {
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

func PrintProjects(projects []config.Project) {
	for _, project := range projects {
		fmt.Println(project.Name)
	}
}
