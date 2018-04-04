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
	labels := []string{}
	project := ""
	sort := "ID"

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
			client.PrintTasks(os.Stdout, tasks, columns, sort)
		},
	}

	listCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")
	listCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "filter tasks with selected task labels")
	listCmd.Flags().StringVarP(&project, "project", "p", project, "project from which to list tasks")
	listCmd.Flags().StringVarP(&sort, "sort", "s", sort, "sort column")

	rootCmd.AddCommand(listCmd)
}
