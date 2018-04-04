package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	columns := []string{
		"ID",
		"Content",
	}
	done := false

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
			client.PrintTasks(os.Stdout, tasks, columns)
			if done {
				// this should be a single item
				for _, t := range tasks {
					rootClient.CloseTask(t)
				}
			}
		},
	}

	addCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")
	addCmd.Flags().BoolVarP(&done, "done", "d", done, "complete the task immediately")

	rootCmd.AddCommand(addCmd)
}
