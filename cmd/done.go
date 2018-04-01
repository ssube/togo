package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	doneCmd := &cobra.Command{
		Use:   "done",
		Short: "mark a task or occurence as completed",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					log.Fatalf("error getting id: %s", err.Error())
				}
				rootClient.CloseTask(client.Task{
					ID: id,
				})
			}
		},
	}

	rootCmd.AddCommand(doneCmd)
}
