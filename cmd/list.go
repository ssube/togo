package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	labels := make([]string, 0)
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
			tasks, _ := rootClient.GetTasks(project, args, labels)
			client.PrintTasks(tasks)
		},
	}

	listCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "filter tasks with selected task labels")
	listCmd.Flags().StringVarP(&project, "project", "p", project, "project from which to list tasks")

	rootCmd.AddCommand(listCmd)
}
