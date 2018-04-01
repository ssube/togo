package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

var rootClient = &client.Client{}
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "togo is a todoist client in go",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("root cmd")
	},
}

func Execute(client *client.Client) {
	rootClient = client
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
