package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type DueDate struct {
	Date     string `json:"date" yaml:"date"`
	DateTime string `json:"datetime" yaml:"datetime"`
	Timezone string `json:"timezone" yaml:"timezone"`
}

// Task model
// https://developer.todoist.com/rest/v8/#tasks
type Task struct {
	Content  string  `json:"content" yaml:"content"`
	Due      DueDate `json:"due" yaml:"due"`
	ID       int     `json:"id" yaml:"id,omitempty"`
	Labels   []int   `json:"label_ids" yaml:"label_ids"`
	Order    int     `json:"order" yaml:"order,omitempty"`
	Priority int     `json:"priority" yaml:"priority,omitempty"`
	Project  int     `json:"project_id" yaml:"project_id"`
}

// PrintTasks in a table
func PrintTasks(f *os.File, tasks []Task, cols []string, sortCol string, dateFmt string) {
	w := CreateTable(f, cols)
	SortByField(tasks, sortCol)

	// prepare a slice for cols and tabs
	for _, t := range tasks {
		fields := GetFields(&t, cols, dateFmt)
		fmt.Fprintln(w, Tabulate(fields)...)
	}

	w.Flush()
}

// ParseTasks from a byte array
func ParseTasks(data []byte) ([]Task, error) {
	out := make([]Task, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

func (c *Client) BuildFilter(required []string, optional []string) string {
	filter := make([]string, 0)

	if len(required) > 0 {
		filter = append(filter, "("+strings.Join(required, " & ")+")")
	}

	if len(optional) > 0 {
		filter = append(filter, "("+strings.Join(optional, " | ")+")")
	}

	return strings.Join(filter, " & ")
}

// GetTasks lists incomplete and recurring tasks
func (c *Client) GetTasks(project string, required []string, optional []string) ([]Task, error) {
	r := c.Request()

	if project != "" {
		r = r.SetQueryParam("project_id", project)
	}

	// build the filters
	r = r.SetQueryParam("filter", c.BuildFilter(required, optional))

	resp, err := r.Get(c.Resolve("tasks"))
	if err != nil {
		log.Printf("error listing tasks: %s", err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("unexpected status listing tasks: %d", resp.StatusCode())
		return nil, errors.New("unexpected status code")
	}

	tasks, err := ParseTasks(resp.Body())
	if err != nil {
		log.Printf("error parsing tasks: %s", err.Error())
		return nil, err
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Order < tasks[j].Order
	})

	return tasks, nil
}

// AddTask creates a new task
func (c *Client) AddTask(task Task) ([]Task, error) {
	post, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("error formatting task: %s", err.Error())
	}

	resp, err := c.Request().
		SetHeader("Content-Type", "application/json").
		SetBody(post).
		Post(c.Resolve("tasks"))
	if err != nil {
		log.Printf("error adding task: %s", err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("response status: %s", resp.Status())
		return nil, errors.New("unexpected response status code")
	}

	body := Task{}
	err = yaml.Unmarshal(resp.Body(), &body)
	if err != nil {
		log.Printf("error parsing response: %s", err.Error())
		return nil, err
	}

	return []Task{
		body,
	}, nil
}

// CloseTask marks an existing task as complete, by ID
func (c *Client) CloseTask(task Task) error {
	if task.ID == 0 {
		log.Fatal("invalid task id")
	}

	log.Printf("closing %d", task.ID)

	path := c.Resolve("tasks", strconv.Itoa(task.ID), "close")
	resp, err := c.Request().Post(path)
	if err != nil {
		log.Fatalf("error adding task: %s", err.Error())
	}

	if resp.StatusCode() != 204 {
		log.Printf("response status: %s", resp.Status())
		return errors.New("unexpected response status code")
	}

	return nil
}
