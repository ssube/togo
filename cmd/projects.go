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

	projectsCmd := &cobra.Command{
		Use:   "projects",
		Short: "list projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects, err := rootClient.GetProjects()
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}

			columns := rootClient.Columns(columns, rootColumns, rootClient.Config().Default.Projects.Columns)
			sort := rootClient.Sort(sort, rootSort, rootClient.Config().Default.Projects.Sort)
			client.PrintProjects(os.Stdout, projects, columns, sort)
		},
	}

	rootCmd.AddCommand(projectsCmd)
}
