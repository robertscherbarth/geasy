package operations

import (
	"context"
	"google.golang.org/api/cloudresourcemanager/v2"
	"google.golang.org/api/cloudresourcemanager/v3"
)

type ResourceManager struct{}

func (r *ResourceManager) ListFolder(orgID string) ([]*cloudresourcemanager.Folder, error) {
	service, err := cloudresourcemanager.NewService(context.Background())
	if err != nil {
		return nil, err
	}

	service := cloudresourcemanager.NewService()
	/*response, err := service.Folders.List().Do()
	if err != nil {
		return nil, err
	}*/

	return response.Folders, nil
}
