package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	columns := []string{
		"ID",
		"Name",
	}
	sort := "ID"

	labelsCmd := &cobra.Command{
		Use:   "labels",
		Short: "list labels",
		Run: func(cmd *cobra.Command, args []string) {
			labels, err := rootClient.GetLabels()
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}
			client.PrintLabels(os.Stdout, labels, columns, sort)
		},
	}

	labelsCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")
	labelsCmd.Flags().StringVarP(&sort, "sort", "s", sort, "sort column")

	rootCmd.AddCommand(labelsCmd)
}
