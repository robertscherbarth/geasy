package operations

import (
	"context"
	"google.golang.org/api/iam/v1"
)

type IAM struct{}

func (i *IAM) ListServiceAccounts(project string) (*iam.ListServiceAccountsResponse, error) {
	service, err := iam.NewService(context.Background())
	if err != nil {
		return nil, err
	}

	response, err := service.Projects.ServiceAccounts.List("projects/" + project).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
