package get

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/robertscherbarth/geasy/internal/operations"
)

func newIAMCmd() *cobra.Command {
	var project string

	cmdIAM := &cobra.Command{
		Use:     "iam",
		Short:   "show current runtime configuration",
		Example: "geasy get runtime --project my-project",
		Run: func(cmd *cobra.Command, args []string) {
			if project == "" {
				cobra.CheckErr(errors.New("project has to be set"))
			}

			service := operations.IAM{}

			list, err := service.ListServiceAccounts(project)
			if err != nil {
				cobra.CheckErr(err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"project", "name", "e-mail", "description"})

			for _, v := range list.Accounts {
				table.Append([]string{project, v.Name, v.Email, v.Description})
			}

			table.Render()
		},
	}

	cmdIAM.Flags().StringVarP(&project, "project", "p", "", "project")

	return cmdIAM
}
