package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		rootClient.AddTask(client.Task{})
	},
}
