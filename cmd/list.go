package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	columns := []string{
		"ID",
		"Content",
	}
	labels := []string{}
	project := ""

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list pending tasks",
		Args:  cobra.RangeArgs(0, 8),
		Run: func(cmd *cobra.Command, args []string) {
			labelTags := make([]string, len(labels))
			for i, l := range labels {
				if strings.HasPrefix(l, "@") {
					labelTags[i] = l
				} else {
					labelTags[i] = "@" + l
				}
			}
			tasks, err := rootClient.GetTasks(project, args, labels)
			if err != nil {
				log.Fatalf("error listing tasks: %s", err.Error())
			}
			client.PrintTasks(tasks, columns)
		},
	}

	listCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")
	listCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "filter tasks with selected task labels")
	listCmd.Flags().StringVarP(&project, "project", "p", project, "project from which to list tasks")

	rootCmd.AddCommand(listCmd)
}
