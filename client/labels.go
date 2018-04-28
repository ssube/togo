package client

import (
	"errors"
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Label model
// https://developer.todoist.com/rest/v8/#labels
type Label struct {
	ID    int    `json:"id" yaml:"id"`
	Name  string `json:"name" yaml:"name"`
	Order int    `json:"order" yaml:"order"`
}

// ParseLabels from a byte array
func ParseLabels(data []byte) ([]Label, error) {
	out := make([]Label, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

// PrintLabels after sorting, with column headers
func PrintLabels(f *os.File, labels []Label, cols []string, sortCol string) {
	w := CreateTable(f, cols)
	SortByField(labels, sortCol)

	for _, l := range labels {
		fields := GetFields(&l, cols)
		fmt.Fprintln(w, Tabulate(fields)...)
	}

	w.Flush()
}

// FindLabel by name
func (c *Client) FindLabel(name string) (Label, error) {
	labels, err := c.GetLabels()
	if err != nil {
		return Label{}, err
	}

	for _, v := range labels {
		if v.Name == name {
			return v, nil
		}
	}

	return Label{}, errors.New("label not found")
}

// GetLabels from API
func (c *Client) GetLabels() ([]Label, error) {
	body, _, err := c.Get("labels")
	if err != nil {
		log.Printf("error getting labels: %s", err.Error())
		return nil, err
	}

	labels, err := ParseLabels(body)
	if err != nil {
		log.Printf("error parsing labels: %s", err.Error())
		return nil, err
	}

	return labels, nil
}
