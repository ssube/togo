package client

import (
	"errors"
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Project struct {
	ID     int    `json:"id" yaml:"id"`
	Indent int    `json:"indent" yaml:"indent"`
	Name   string `json:"name" yaml:"name"`
	Order  int    `json:"order" yaml:"order"`
}

func ParseProjects(data []byte) ([]Project, error) {
	out := make([]Project, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

func PrintProjects(f *os.File, projects []Project, cols []string, sortCol string) {
	w := PrintTable(f, cols)
	SortField(projects, sortCol)

	// prepare a slice for cols and tabs
	for _, p := range projects {
		fields := GetFields(&p, cols)
		fmt.Fprintln(w, Tabulate(fields)...)
	}

	w.Flush()
}

func (c *Client) GetProjects() ([]Project, error) {
	resp, err := c.Request().Get(c.GetEndpoint("projects"))
	if err != nil {
		log.Printf("error getting projects: %s", err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("unexpected response status: %d", resp.StatusCode())
		return nil, errors.New("unexpected response status")
	}

	projects, err := ParseProjects(resp.Body())
	if err != nil {
		log.Printf("error parsing projects: %s", err.Error())
		return nil, err
	}

	return projects, nil
}
