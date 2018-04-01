package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := rootClient.GetTasks()
		client.PrintTasks(tasks)
	},
}
