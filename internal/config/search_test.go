package config

import "testing"

type Fixtures struct {
	projects []*Project
	cfg      *Config
}

func fixtures() Fixtures {

	projects := []*Project{
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

	return Fixtures{
		projects,
		&cfg,
	}

}

func TestSearchProjects(t *testing.T) {
	cfg := fixtures().cfg

	var tests = []struct {
		input    string
		expected int
	}{
		{"xxx", 0},
		{"project", 1},
		{"o", 2},
	}

	for _, test := range tests {
		projects := SearchProjects(test.input, *cfg)
		if len(projects) != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}",
				test.input, test.expected, projects)
		}
	}

}

func TestMatchProjectsMatch(t *testing.T) {
	cfg := fixtures().cfg

	project := MatchProjects("project", *cfg)
	if project.Name != "project" {
		t.Error("Test Failed: {} inputed, {} expected, recieved: {}",
			"project", "project", project.Name)
	}
}

func TestMatchProjectsNoMatch(t *testing.T) {
	cfg := fixtures().cfg

	project := MatchProjects("xxx", *cfg)
	if project != nil {
		t.Error("Test Failed: {} inputed, {} expected, recieved: {}",
			"xxx", "no match", project)
	}
}

func TestGetProjects(t *testing.T) {
	cfg := fixtures().cfg

	var tests = []struct {
		input  string
		hasErr bool
	}{
		{"xxx", true},
		{"project", false},
		{"o", true},
	}

	for _, test := range tests {
		project, err := GetOneProject(test.input, *cfg)
		hasErr := err != nil
		if test.hasErr != hasErr {
			t.Error("Test Failed: {} inputed, {} expected, received: {}",
				test.input, test.hasErr, hasErr)
		}
		if err != nil {
			if project != nil {
				t.Error("Test Failed: input: {}, expected: {}, received: {}",
					test.input, "project is null", project)
			}
		}
	}
}
