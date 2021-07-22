package operations

import (
	"context"
	"fmt"
	"github.com/robertscherbarth/geasy/internal/printer"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

type RuntimeConfig struct{}

func (r *RuntimeConfig) GetVariablesFromProject(project string) ([]*runtimeconfig.Variable, error) {
	ctx := context.Background()
	runtimeconfigService, err := runtimeconfig.NewService(ctx)
	if err != nil {
		return nil, err
	}

	configListCall := runtimeconfigService.Projects.Configs.List(fmt.Sprintf("projects/%s", project))
	configsResponse, err := configListCall.Do()
	if err != nil {
		return nil, err
	}

	variables := make([]*runtimeconfig.Variable, 0)
	for _, v := range configsResponse.Configs {
		call := runtimeconfigService.Projects.Configs.Variables.List(v.Name).ReturnValues(true)
		list, err := call.Do()
		if err != nil {
			printer.Errorf("could not add variables cause of : %v", err)
			continue
		}

		variables = append(variables, list.Variables...)
	}

	return variables, nil
}
