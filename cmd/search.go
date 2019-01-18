package cmd

import (
	"fmt"
	"strings"
)

func findProject(query string, config Config) Project {
	var projects []Project
	for _, project := range config.Projects {
		if strings.Contains(project.Name, query) {
			// if query == project.Name {
			projects = append(projects, project)
		}
	}
	if len(projects) == 0 {
		Terminate("Project not found")
		fmt.Println("got here")
	} else if len(projects) > 1 {
		for _, project := range projects {
			fmt.Println(project.Name)
		}
		Terminate("\nMore than one match.")
	}
	return projects[0]
}
