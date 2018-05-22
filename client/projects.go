package client

import (
	"errors"
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Project model
// https://developer.todoist.com/rest/v8/#projects
type Project struct {
	ID     int    `json:"id" yaml:"id"`
	Indent int    `json:"indent" yaml:"indent"`
	Name   string `json:"name" yaml:"name"`
	Order  int    `json:"order" yaml:"order"`
}

// ParseProjects from byte array
func ParseProjects(data []byte) ([]Project, error) {
	out := make([]Project, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

// PrintProjects after sorting, with column headers
func PrintProjects(f *os.File, projects []Project, cols []string, sortCol string, dateFmt string) {
	w := CreateTable(f, cols)
	SortByField(projects, sortCol)

	// prepare a slice for cols and tabs
	for _, p := range projects {
		fields := GetFields(&p, cols, dateFmt)
		fmt.Fprintln(w, Tabulate(fields)...)
	}

	w.Flush()
}

// FindProject by name
func (c *Client) FindProject(name string) (Project, error) {
	projects, err := c.GetProjects()
	if err != nil {
		return Project{}, err
	}

	for _, p := range projects {
		if p.Name == name {
			return p, nil
		}
	}

	return Project{}, errors.New("project not found")
}

// GetProjects from API
func (c *Client) GetProjects() ([]Project, error) {
	body, _, err := c.Get("projects")
	if err != nil {
		log.Printf("error getting projects: %s", err.Error())
		return nil, err
	}

	projects, err := ParseProjects(body)
	if err != nil {
		log.Printf("error parsing projects: %s", err.Error())
		return nil, err
	}

	return projects, nil
}
