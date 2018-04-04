package client

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

func PrintProjects(projects []Project) {
	w := tabwriter.NewWriter(os.Stdout, 4, 2, 2, ' ', 0)
	fmt.Fprintln(w, "ID", "\t", "Name")

	// prepare a slice for cols and tabs
	for _, p := range projects {
		fmt.Fprintln(w, p.ID, "\t", p.Name)
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
