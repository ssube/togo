package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// Task model from API
type Task struct {
	Content  string `json:"content" yaml:"content"`
	ID       int    `json:"id" yaml:"id,omitempty"`
	Order    int    `json:"order" yaml:"order,omitempty"`
	Priority int    `json:"priority" yaml:"priority,omitempty"`
}

// PrintTasks in a table
func PrintTasks(f *os.File, tasks []Task, cols []string, sortCol string) {
	w := PrintTable(f, cols)
	SortField(tasks, sortCol)

	// prepare a slice for cols and tabs
	taskCols := make([]string, len(cols))
	for _, t := range tasks {
		vc := reflect.ValueOf(&t)

		for i, c := range cols {
			field := vc.Elem().FieldByName(c)

			if !field.IsValid() {
				log.Fatalf("missing column: %s", c)
			}

			ftype := field.Type()
			fkind := ftype.Kind()

			switch fkind {
			case reflect.Int:
				taskCols[i] = strconv.FormatInt(field.Int(), 10)
			case reflect.String:
				taskCols[i] = field.String()
			default:
				taskCols[i] = "."
			}
		}
		fmt.Fprintln(w, Tabulate(taskCols)...)
	}

	w.Flush()
}

// Parse a reponse into list of tasks
func ParseTasks(data []byte) ([]Task, error) {
	out := make([]Task, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

// GetTasks lists incomplete and recurring tasks
func (c *Client) GetTasks(project string, required []string, optionals []string) ([]Task, error) {
	r := c.Request()

	if project != "" {
		r = r.SetQueryParam("project_id", project)
	}

	// build the filters
	filter := make([]string, 0)

	if len(required) > 0 {
		filter = append(filter, "("+strings.Join(required, " & ")+")")
	}

	if len(optionals) > 0 {
		filter = append(filter, "("+strings.Join(optionals, " | ")+")")
	}

	r = r.SetQueryParam("filter", strings.Join(filter, " & "))

	resp, err := r.Get(c.GetEndpoint("tasks"))
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
		Post(c.GetEndpoint("tasks"))
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

	path := c.GetEndpoint("tasks", strconv.Itoa(task.ID), "close")
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
