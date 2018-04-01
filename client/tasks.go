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
	"text/tabwriter"

	"github.com/ssube/togo/config"
	"gopkg.in/resty.v1"
	"gopkg.in/yaml.v2"
)

type Client struct {
	client *resty.Client
	config *config.Config
	root   string
}

type Task struct {
	Content  string `json:"content" yaml:"content"`
	ID       int    `json:"id" yaml:"id,omitempty"`
	Order    int    `json:"order" yaml:"order,omitempty"`
	Priority int    `json:"priority" yaml:"priority,omitempty"`
}

func PrintTasks(tasks []Task) {
	w := tabwriter.NewWriter(os.Stdout, 4, 2, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "id", "\t", "priority", "\t", "content")
	for _, t := range tasks {
		fmt.Fprintln(w, t.ID, "\t", t.Priority, "\t", t.Content)
	}
	w.Flush()
}

func New(config *config.Config) *Client {
	client := &Client{
		client: resty.New(),
		config: config,
	}
	return client
}

func (c *Client) Endpoint(parts ...string) string {
	path := []string{
		c.config.Root,
	}

	return strings.Join(append(path, parts...), "/")
}

func (c *Client) Request() *resty.Request {
	return c.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
}

func (c *Client) Parse(data []byte) ([]Task, error) {
	out := make([]Task, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

func (c *Client) GetTasks() ([]Task, error) {
	resp, err := c.Request().Get(c.Endpoint("tasks"))
	if err != nil {
		log.Printf("error listing tasks: %s", err.Error())
		return nil, err
	}

	tasks, err := c.Parse(resp.Body())
	if err != nil {
		log.Printf("error parsing tasks: %s", err.Error())
		return nil, err
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Order < tasks[j].Order
	})

	return tasks, nil
}

func (c *Client) AddTask(task Task) ([]Task, error) {
	post, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("error formatting task: %s", err.Error())
	}

	resp, err := c.Request().
		SetHeader("Content-Type", "application/json").
		SetBody(post).
		Post(c.Endpoint("tasks"))
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

func (c *Client) CloseTask(task Task) error {
	if task.ID == 0 {
		log.Fatal("invalid task id")
	}

	log.Printf("closing %d", task.ID)

	resp, err := c.Request().Post(c.Endpoint("tasks", strconv.Itoa(task.ID), "close"))
	if err != nil {
		log.Fatalf("error adding task: %s", err.Error())
	}

	if resp.StatusCode() != 204 {
		log.Printf("response status: %s", resp.Status())
		return errors.New("unexpected response status code")
	}

	return nil
}
