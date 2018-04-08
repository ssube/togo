package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

var rootFilter = "overdue | today"
var rootClient = &client.Client{}
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "togo is a todoist client in go",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := rootClient.GetTasks("", []string{
			rootFilter,
		}, []string{})
		fmt.Printf("%d tasks to go\n", len(tasks))
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

func init() {
	rootCmd.Flags().StringVarP(&rootFilter, "filter", "f", rootFilter, "list filter")
}
