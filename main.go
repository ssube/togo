package main

import (
	"github.com/ssube/togo/client"
	"github.com/ssube/togo/cmd"
	"github.com/ssube/togo/config"
)

func main() {
	config := config.New(".togo.yml")
	client := client.New(config)
	client.GetTasks()

	cmd.Execute()
}
