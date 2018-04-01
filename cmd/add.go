package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	content := ""

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			rootClient.AddTask(client.Task{
				Content: content,
			})
		},
	}

	addCmd.Flags().StringVarP(&content, "content", "c", content, "task content")

	rootCmd.AddCommand(addCmd)
}
