package client

import (
	"os"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	tasks := []Task{
		Task{
			ID:      1,
			Content: "test",
		},
		Task{
			ID:      2,
			Content: "testing",
		},
	}

	PrintTasks(os.Stdout, tasks, []string{
		"Content",
	}, "Content", "2006-01-02")
}

func TestParseTasks(t *testing.T) {
	data := `
- id: 0
  name: test
`

	tasks, err := ParseTasks([]byte(data))

	if err != nil {
		t.Fatalf("error parsing tasks: %s", err.Error())
	}

	if len(tasks) != 1 {
		t.Fatalf("wrong number of tasks: %d", len(tasks))
	}
}
