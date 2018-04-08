package cmd

import (
	"log"
	"os"
	"strconv"
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
	labels := []string{}
	project := ""
	sort := "ID"

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := client.Task{
				Content: strings.Join(args, " "),
			}

			// resolve labels
			if len(labels) > 0 {
				labelIDs := make([]int, len(labels))
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

				task.Labels = labelIDs
			}

			// resolve project
			if project != "" {
				project, err := rootClient.FindProject(project)
				if err != nil {
					log.Fatalf("error getting projects: %s", err.Error())
				}

				task.Project = project.ID
			}

			// add task
			tasks, err := rootClient.AddTask(task)
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}

			client.PrintTasks(os.Stdout, tasks, columns, sort)
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
	addCmd.Flags().StringSliceVarP(&labels, "labels", "l", labels, "task labels")
	addCmd.Flags().StringVarP(&project, "project", "p", project, "task project")
	addCmd.Flags().StringVarP(&sort, "sort", "s", sort, "sort column")

	rootCmd.AddCommand(addCmd)
}
