package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := rootClient.GetTasks()

		w := tabwriter.NewWriter(os.Stdout, 4, 2, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintln(w, "id", "\t", "priority", "\t", "content")
		for _, t := range tasks {
			fmt.Fprintln(w, t.ID, "\t", t.Priority, "\t", t.Content)
		}
		w.Flush()
	},
}
