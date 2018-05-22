package cmd

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	columns := []string{
		"ID",
		"Content",
	}
	date := time.RFC3339
	labels := rootConfig.Default.Tasks.Labels
	project := rootConfig.Default.Tasks.Project
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

			columns := rootClient.Columns(columns, rootColumns, rootClient.Config().Default.Tasks.Columns)
			date := rootClient.Sort(date, rootDate, rootClient.Config().Default.Date)
			sort := rootClient.Sort(sort, rootSort, rootClient.Config().Default.Tasks.Sort)
			client.PrintTasks(os.Stdout, tasks, columns, sort, date)
		},
	}

	listCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "filter tasks with selected task labels")
	listCmd.Flags().StringVarP(&project, "project", "p", project, "project from which to list tasks")

	rootCmd.AddCommand(listCmd)
}
