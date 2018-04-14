package main

import (
	"log"

	"github.com/ssube/togo/client"
	"github.com/ssube/togo/cmd"
	"github.com/ssube/togo/config"
)

func main() {
	userConfig := config.New(".togo.yml")
	log.SetFlags(0)

	client := client.New(userConfig)
	cmd.Execute(client)
}
