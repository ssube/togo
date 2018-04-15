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

// persistent flags
var rootColumns = []string{}
var rootSort = ""

// flags
var rootFilter = ""

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "togo is a todoist client in go",
	Run: func(cmd *cobra.Command, args []string) {
		filter := rootClient.Sort("overdue | today", rootFilter, rootClient.Config().Default.Root.Filter)
		tasks, _ := rootClient.GetTasks("", []string{
			filter,
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
	rootCmd.PersistentFlags().StringVarP(&rootSort, "sort", "s", "", "sort column")

	// root-only flags
	rootCmd.Flags().StringVarP(&rootFilter, "filter", "f", rootFilter, "list filter")
}
