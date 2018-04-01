package client

import (
	"fmt"
	"log"

	"github.com/ssube/togo/config"
	"gopkg.in/resty.v1"
	"gopkg.in/yaml.v2"
)

type Client struct {
	client *resty.Client
	config *config.Config
	root   string
}

func New(config *config.Config) *Client {
	client := &Client{
		client: resty.New(),
		config: config,
		root:   "https://beta.todoist.com/API/v8/",
	}
	return client
}

func (c *Client) Endpoint(path string) string {
	return fmt.Sprintf("%s/%s", c.root, path)
}

func (c *Client) Request() *resty.Request {
	return c.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
}

func (c *Client) Parse(data []byte) ([]map[string]interface{}, error) {
	out := make([]map[string]interface{}, 0)
	err := yaml.Unmarshal(data, &out)

	return out, err
}

func (c *Client) GetTasks() {
	resp, err := c.Request().Get(c.Endpoint("tasks"))
	if err != nil {
		log.Fatalf("error listing tasks: %s", err.Error())
	}

	log.Printf("status: %s", resp.Status())

	body, err := c.Parse(resp.Body())
	if err != nil {
		log.Fatalf("error parsing tasks: %s", err.Error())
	}

	for _, task := range body {
		log.Printf("%t: %s", task["completed"], task["content"])
	}
}

func (c *Client) AddTask() {
	resp, err := c.Request().Post(c.Endpoint("tasks"))
	if err != nil {
		log.Fatalf("error adding task: %s", err.Error())
	}

	body, err := c.Parse(resp.Body())
	if err != nil {
		log.Fatalf("error parsing tasks: %s", err.Error())
	}

	log.Printf("status: %s", resp.Status())
	log.Printf("body: %s", body)
}
