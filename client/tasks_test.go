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
	}, "Content")
}
