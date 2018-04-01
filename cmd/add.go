package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	columns := []string{}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := rootClient.AddTask(client.Task{
				Content: strings.Join(args, " "),
			})
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}
			client.PrintTasks(tasks, columns)
		},
	}

	addCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")

	rootCmd.AddCommand(addCmd)
}
