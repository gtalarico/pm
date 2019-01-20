package config

import "testing"

func TestSearchProjects(t *testing.T) {
	projects := []Project{
		{
			Name: "project",
			Path: "~",
		},
		{
			Name: "another",
			Path: "~",
		},
	}
	cfg := Config{projects}

	var tests = []struct {
		input    string
		expected int
	}{
		{"xxx", 0},
		{"project", 1},
		{"o", 2},
	}

	for _, test := range tests {
		projects := SearchProjects(test.input, cfg)
		if len(projects) != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, projects)
		}
	}

}
