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

	projectsCmd := &cobra.Command{
		Use:   "projects",
		Short: "list projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects, err := rootClient.GetProjects()
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}
			client.PrintProjects(os.Stdout, projects, columns, sort)
		},
	}

	projectsCmd.Flags().StringSliceVarP(&columns, "columns", "c", columns, "display columns")
	projectsCmd.Flags().StringVarP(&sort, "sort", "s", sort, "sort column")

	rootCmd.AddCommand(projectsCmd)
}
