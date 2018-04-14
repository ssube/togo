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
	sort := "Name"

	labelsCmd := &cobra.Command{
		Use:   "labels",
		Short: "list labels",
		Run: func(cmd *cobra.Command, args []string) {
			labels, err := rootClient.GetLabels()
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}

			columns := rootClient.Columns(columns, rootColumns, rootClient.Config().Default.Labels.Columns)
			sort := rootClient.Sort(sort, rootSort, rootClient.Config().Default.Labels.Sort)
			client.PrintLabels(os.Stdout, labels, columns, sort)
		},
	}

	rootCmd.AddCommand(labelsCmd)
}
