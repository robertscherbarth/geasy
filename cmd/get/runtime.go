package get

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

const variables = "variables"

func newRuntimeConfigurationCmd() *cobra.Command {
	var project string
	var configs string

	cmdRuntime := &cobra.Command{
		Use:     "runtime",
		Short:   "show current runtime configuration",
		Example: "geasy get runtime --project my-project",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			runtimeconfigService, err := runtimeconfig.NewService(ctx)
			cobra.CheckErr(err)

			if project == "" {
				cobra.CheckErr(errors.New("project has to be set"))
			}

			//TODO: move this to operations in a clean client

			/*var currentProject, currentConfig, currentVariables string
			runtimeconfigService.Projects.Configs.Variables.Get(fmt.Sprintf("projects/%s/configs/%s/variables/%s", currentProject, currentConfig, currentVariables))
			*/

			configListCall := runtimeconfigService.Projects.Configs.List(fmt.Sprintf("projects/%s", project))
			configsResponse, err := configListCall.Do()
			cobra.CheckErr(err)

			variables := make([]*runtimeconfig.Variable, 0)
			for _, v := range configsResponse.Configs {
				call := runtimeconfigService.Projects.Configs.Variables.List(v.Name).ReturnValues(true)
				list, err := call.Do()
				cobra.CheckErr(err)

				variables = append(variables, list.Variables...)
			}
			/*
				printer.Infof("current project %v", project)
				parent := fmt.Sprintf("projects/%s/configs/%s", project, configs)
				call := runtimeconfigService.Projects.Configs.Variables.List(parent).ReturnValues(true)
				list, err := call.Do()
				cobra.CheckErr(err)*/

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"project", "key", "value", "update-time"})

			for _, v := range variables {
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
