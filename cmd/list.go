package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("list tasks")
		// client.GetTasks()
	},
}
