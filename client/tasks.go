package client

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

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

func New(config *config.Config) *Client {
	client := &Client{
		client: resty.New(),
		config: config,
		root:   "https://beta.todoist.com/API/v8",
	}
	return client
}

func (c *Client) Endpoint(parts ...string) string {
	path := []string{
		c.root,
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
		log.Fatalf("error listing tasks: %s", err.Error())
	}

	tasks, err := c.Parse(resp.Body())
	if err != nil {
		log.Fatalf("error parsing tasks: %s", err.Error())
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Order < tasks[j].Order
	})

	return tasks, nil
}

func (c *Client) AddTask(task Task) error {
	post, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("error formatting task: %s", err.Error())
	}

	log.Printf("adding: %s", post)

	resp, err := c.Request().
		SetHeader("Content-Type", "application/json").
		SetBody(post).
		Post(c.Endpoint("tasks"))
	if err != nil {
		log.Fatalf("error adding task: %s", err.Error())
	}

	log.Printf("status: %s", resp.Status())
	log.Printf("body: %s", string(resp.Body()))

	_, err = c.Parse(resp.Body())
	if err != nil {
		log.Fatalf("error parsing tasks: %s", err.Error())
	}

	return nil
}

func (c *Client) CloseTask(task Task) error {
	if task.ID == 0 {
		log.Fatal("invalid task id")
	}

	path := c.Endpoint("tasks", strconv.Itoa(task.ID), "close")
	log.Printf("closing: %s", path)

	resp, err := c.Request().Post(path)
	if err != nil {
		log.Fatalf("error adding task: %s", err.Error())
	}

	log.Printf("status: %s", resp.Status())
	log.Printf("body: %s", resp.Body())

	return nil
}
