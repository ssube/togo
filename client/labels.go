package client

import (
	"errors"
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Label struct {
	ID    int    `json:"id" yaml:"id"`
	Name  string `json:"name" yaml:"name"`
	Order int    `json:"order" yaml:"order"`
}

func ParseLabels(data []byte) ([]Label, error) {
	out := make([]Label, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

func PrintLabels(f *os.File, labels []Label, cols []string, sortCol string) {
	w := PrintTable(f, cols)
	SortField(labels, sortCol)

	for _, l := range labels {
		fields := GetFields(&l, cols)
		fmt.Fprintln(w, Tabulate(fields)...)
	}

	w.Flush()
}

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

func (c *Client) GetLabels() ([]Label, error) {
	resp, err := c.Request().Get(c.GetEndpoint("labels"))
	if err != nil {
		log.Printf("error getting labels: %s", err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("unexpected response status: %d", resp.StatusCode())
		return nil, errors.New("unexpected response status")
	}

	labels, err := ParseLabels(resp.Body())
	if err != nil {
		log.Printf("error parsing labels: %s", err.Error())
		return nil, err
	}

	return labels, nil
}
