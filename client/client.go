package client

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

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

func PrintTable(f *os.File, cols []string) *tabwriter.Writer {
	w := tabwriter.NewWriter(f, 4, 2, 2, ' ', 0)
	fmt.Fprintln(w, Tabulate(cols)...)
	return w
}

func GetFields(val interface{}, cols []string) []string {
	out := make([]string, len(cols))
	rv := reflect.ValueOf(val).Elem()

	for i, c := range cols {
		field := rv.FieldByName(c)

		if !field.IsValid() {
			log.Fatalf("missing column: %s", c)
		}

		ftype := field.Type()
		fkind := ftype.Kind()

		switch fkind {
		case reflect.Int:
			out[i] = strconv.FormatInt(field.Int(), 10)
		case reflect.String:
			out[i] = field.String()
		case reflect.Slice:
			out[i] = fmt.Sprintf("%v", field)
		default:
			log.Printf("unknown type for column %s: %v", c, ftype)
			out[i] = "."
		}
	}

	return out
}

func SortField(vals interface{}, col string) {
	vv := reflect.ValueOf(vals)
	head := vv.Index(0)
	sortField := head.FieldByName(col)

	if !sortField.IsValid() {
		log.Fatalf("missing sort column: %s", col)
	}

	sort.Slice(vals, func(i, j int) bool {
		it := vv.Index(i)
		jt := vv.Index(j)

		is := it.FieldByName(col).String()
		js := jt.FieldByName(col).String()

		return is < js
	})
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
