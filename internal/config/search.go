package config

import (
	"errors"
	"fmt"
	"strings"
)

func SearchProjects(query string, config Config) []Project {
	var projects []Project
	for _, project := range config.Projects {
		if strings.Contains(project.Name, query) {
			projects = append(projects, project)
		}
	}
	return projects
}

func MatchProjects(query string, config Config) (project Project) {
	for _, p := range config.Projects {
		if p.Name == query {
			project = p
		}
	}
	return
}

func GetOneProject(query string, config Config) (project Project, err error) {
	projects := SearchProjects(query, config)
	numProjects := len(projects)

	if numProjects == 1 {
		project = projects[0]
	}
	if numProjects == 0 {
		err = errors.New("no match")
	}
	if numProjects > 1 {
		exactMatch := MatchProjects(query, config)
		if exactMatch.Name == query {
			project = exactMatch
		} else {
			for _, project := range projects {
				fmt.Println(project.Name)
			}
			err = errors.New("multiple matches")
		}
	}
	return
}
