package client

import (
	"fmt"
	"strings"

	"github.com/ssube/togo/config"
	"gopkg.in/resty.v1"
)

// Client for API
type Client struct {
	client *resty.Client
	config *config.Config
}

func Tabulate(cols []string) []interface{} {
	tabs := make([]interface{}, len(cols)*2)
	for i, c := range cols {
		tabs[i*2] = c
		tabs[(i*2)+1] = "\t"
	}
	return tabs
}

// New client
func New(config *config.Config) *Client {
	client := &Client{
		client: resty.New(),
		config: config,
	}
	return client
}

// GetEndpoint formats a partial path as an API endpoint
func (c *Client) GetEndpoint(parts ...string) string {
	path := []string{
		c.config.Root,
	}

	return strings.Join(append(path, parts...), "/")
}

// Request an API endpoint with authorization
func (c *Client) Request() *resty.Request {
	return c.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
}
