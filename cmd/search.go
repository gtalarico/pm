package cmd

import (
	"errors"
	"fmt"
	"strings"
)

func searchProjects(query string, config Config, exact bool) []Project {
	var projects []Project
	for _, project := range config.Projects {
		if strings.Contains(project.Name, query) {
			projects = append(projects, project)
		}
	}
	return projects
}

func getOneProject(query string, config Config, exact bool) (project Project, err error) {
	projects := searchProjects(query, config, exact)
	numProjects := len(projects)

	if numProjects == 1 {
		project = projects[0]
	}
	if numProjects == 0 {
		err = errors.New("no match")
	}
	if numProjects > 1 {
		for _, project := range projects {
			fmt.Println(project.Name)
		}
		err = errors.New("multiple matches")
	}
	return
}
