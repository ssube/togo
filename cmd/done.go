package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "mark a task or occurence as completed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do task")
	},
}
