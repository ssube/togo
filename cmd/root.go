package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
	"github.com/ssube/togo/config"
)

// client/config
// TODO: make this less global :(
var rootClient = &client.Client{}
var rootConfig = &config.Config{}

// flags
var rootColumns = []string{}
var rootFilter = "overdue | today"
var rootSort = ""

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
	rootConfig = client.Config()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&rootColumns, "columns", "c", []string{}, "display columns")
	rootCmd.PersistentFlags().StringVarP(&rootFilter, "filter", "f", rootFilter, "list filter")
	rootCmd.PersistentFlags().StringVarP(&rootSort, "sort", "s", "", "sort column")
}
