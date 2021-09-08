package get

import (
	"github.com/olekukonko/tablewriter"
	"github.com/robertscherbarth/geasy/internal/operations"
	"github.com/spf13/cobra"
	"os"
)

func newFolder() *cobra.Command {
	var project string

	cmdIAM := &cobra.Command{
		Use:     "folder",
		Short:   "show folder",
		Example: "geasy get folder --project my-project",
		Run: func(cmd *cobra.Command, args []string) {

			service := operations.ResourceManager{}

			list, err := service.ListFolder("")
			if err != nil {
				cobra.CheckErr(err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"project", "name", "display name", "state"})

			for _, v := range list {
				table.Append([]string{project, v.Name, v.DisplayName, v.State})
			}

			table.Render()
		},
	}

	return cmdIAM
}
