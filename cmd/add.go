package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func resolveLabels(labels []string) []int {
	labelIDs := make([]int, len(labels))

	if len(labels) == 0 {
		return labelIDs
	}

	for i, l := range labels {
		id, err := strconv.Atoi(l)
		if err == nil {
			labelIDs[i] = id
			continue
		}

		label, err := rootClient.FindLabel(l)
		if err != nil {
			log.Fatalf("error getting labels: %s", err.Error())
		}

		labelIDs[i] = label.ID
	}

	return labelIDs
}

func resolveProject(ref string) int {
	if ref == "" {
		return 0
	}

	project, err := rootClient.FindProject(ref)
	if err != nil {
		log.Fatalf("error getting projects: %s", err.Error())
	}

	return project.ID
}

func init() {
	columns := []string{
		"ID",
		"Content",
	}
	done := false
	labels := rootConfig.Default.Tasks.Labels
	project := rootConfig.Default.Tasks.Project
	sort := "ID"

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := client.Task{
				Content: strings.Join(args, " "),
				Labels:  resolveLabels(labels),
				Project: resolveProject(project),
			}

			// add task
			tasks, err := rootClient.AddTask(task)
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}

			columns := rootClient.Columns(columns, rootColumns, rootClient.Config().Default.Projects.Columns)
			sort := rootClient.Sort(sort, rootSort, rootClient.Config().Default.Projects.Sort)
			client.PrintTasks(os.Stdout, tasks, columns, sort)
			if done {
				// this should be a single item
				for _, t := range tasks {
					rootClient.CloseTask(t)
				}
			}
		},
	}

	addCmd.Flags().BoolVarP(&done, "done", "d", done, "complete the task immediately")
	addCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "task labels")
	addCmd.Flags().StringVarP(&project, "project", "p", project, "task project")

	rootCmd.AddCommand(addCmd)
}
