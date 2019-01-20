package commands

import (
	"testing"
)

func TestGetCommand(t *testing.T) {

	// Test Valid Commands
	inputs := [...]string{"go", "list", "add", "remove"}
	for _, input := range inputs {
		command, _ := GetCommand(input)
		if command.Name != input {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, input, command.Name)
		}
	}

	// Test Invalid Command
	input := "invalidcmd"
	_, err := GetCommand(input)
	if err == nil {
		t.Error("Test Failed: {} inputted, {} expected", input, "an error")
	}

}
