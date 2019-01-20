package cli

import "testing"

func TestValidateArgsValid(t *testing.T) {

	// Test Invalid Command
	input := []string{"list"}
	expected := "list"
	cmdName, _, _ := ValidateArgs(input)
	if cmdName != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, cmdName)
	}
}

func TestValidateArgsValidInvalid(t *testing.T) {
	input := []string{"xxx"}
	_, _, err := ValidateArgs(input)
	if err != nil {
		t.Error("Test Failed: {} inputted, {} expected, {} received", input, "error", err)
	}

}
