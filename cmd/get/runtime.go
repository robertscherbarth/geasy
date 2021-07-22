package get

import (
	"errors"
	"fmt"
	"github.com/robertscherbarth/geasy/internal/operations"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newRuntimeConfigurationCmd() *cobra.Command {
	var project string
	var configs string

	cmdRuntime := &cobra.Command{
		Use:     "runtime",
		Short:   "show current runtime configuration",
		Example: "geasy get runtime --project my-project",
		Run: func(cmd *cobra.Command, args []string) {
			if project == "" {
				cobra.CheckErr(errors.New("project has to be set"))
			}

			client := operations.RuntimeConfig{}
			list, err := client.GetVariablesFromProject(project)
			if err != nil {
				cobra.CheckErr(err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"project", "key", "value", "update-time"})

			for _, v := range list {
				name := strings.TrimPrefix(v.Name, fmt.Sprintf("projects/%s/configs/%s/variables/", project, configs))
				name = strings.ReplaceAll(name, "/", ".")
				table.Append([]string{project, name, v.Text, v.UpdateTime})
			}

			table.Render()
		},
	}

	cmdRuntime.Flags().StringVarP(&project, "project", "p", "", "project")
	cmdRuntime.Flags().StringVarP(&configs, "configs", "c", "", "configs")

	return cmdRuntime
}
