package client

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ssube/togo/config"
	"gopkg.in/resty.v1"
)

// Client for API
type Client struct {
	client *resty.Client
	config *config.Config
}

// Tabulate an array of column names by inserting tabs between each
// and converting to a Printf-compatible slice
func Tabulate(cols []string) []interface{} {
	tabs := make([]interface{}, len(cols)*2)
	for i, c := range cols {
		tabs[i*2] = c
		tabs[(i*2)+1] = "\t"
	}
	return tabs
}

// CreateTable with column headers and return a row writer
func CreateTable(f *os.File, cols []string) *tabwriter.Writer {
	w := tabwriter.NewWriter(f, 4, 2, 2, ' ', 0)
	fmt.Fprintln(w, Tabulate(cols)...)
	return w
}

func FormatDate(d DueDate, dateFmt string) (string, error) {
	// if no timezone is provided, this is a simple date (day)
	if d.Timezone == "" {
		return d.Date, nil
	}

	zone, err := time.LoadLocation(d.Timezone)
	if err != nil {
		return d.DateTime, err
	}

	local, err := time.ParseInLocation(time.RFC3339, d.DateTime, zone)
	if err != nil {
		return d.DateTime, err
	}

	return local.Format(dateFmt), nil
}

// GetFields from a column list using reflection
func GetFields(val interface{}, cols []string, dateFmt string) []string {
	out := make([]string, len(cols))
	rv := reflect.ValueOf(val).Elem()

	for i, c := range cols {
		field := rv.FieldByName(c)

		if !field.IsValid() {
			log.Fatalf("missing column: %s", c)
		}

		fieldType := field.Type()
		fieldKind := fieldType.Kind()

		switch fieldKind {
		case reflect.Int:
			out[i] = strconv.FormatInt(field.Int(), 10)
		case reflect.String:
			out[i] = field.String()
		case reflect.Struct:
			val := field.Interface()
			switch val.(type) {
			case DueDate:
				fmt, err := FormatDate(val.(DueDate), dateFmt)
				if err != nil {
					log.Fatalf("unable to format date: %s", err.Error())
				}
				out[i] = fmt
			default:
				log.Printf("unknown struct for column %s: %v", c, val)
			}
		case reflect.Slice:
			fallthrough
		default:
			log.Printf("unknown type for column %s: %v", c, fieldType)
			out[i] = fmt.Sprintf("%v", field)
		}
	}

	return out
}

// SortByField sorts the given slice of rows by the selected column
func SortByField(rows interface{}, column string) {
	vv := reflect.ValueOf(rows)
	head := vv.Index(0)
	SortByField := head.FieldByName(column)

	if !SortByField.IsValid() {
		log.Fatalf("missing sort column: %s", column)
	}

	sort.Slice(rows, func(i, j int) bool {
		it := vv.Index(i)
		jt := vv.Index(j)

		is := it.FieldByName(column).String()
		js := jt.FieldByName(column).String()

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

func (c *Client) Get(parts ...string) (body []byte, status int, err error) {
	resp, err := c.Request().Get(c.Resolve(parts...))
	if err != nil {
		log.Printf("error getting labels: %s", err.Error())
		return
	}

	status = resp.StatusCode()
	if status != 200 {
		log.Printf("unexpected response status: %d", status)
		return nil, status, errors.New("unexpected response status")
	}

	body = resp.Body()
	return
}

// Resolve a partial path as an API endpoint
func (c *Client) Resolve(parts ...string) string {
	path := []string{
		c.config.Root,
	}

	return strings.Join(append(path, parts...), "/")
}

// Request an API endpoint with authorization
func (c *Client) Request() *resty.Request {
	return c.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
}

// Config used by this client
func (c *Client) Config() *config.Config {
	return c.config
}

// Columns to display based on command, config, and flags
func (c *Client) Columns(cmdColumns []string, rootColumns []string, configColumns []string) []string {
	if len(rootColumns) > 0 {
		return rootColumns
	} else if len(configColumns) > 0 {
		return configColumns
	}
	return cmdColumns
}

// Sort column based on command, config, and flags
func (c *Client) Sort(cmdSort string, rootSort string, configSort string) string {
	if len(rootSort) > 0 {
		return rootSort
	} else if len(configSort) > 0 {
		return configSort
	}
	return cmdSort
}
