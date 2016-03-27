package docker

import (
	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
)

// List lists the containers managed by kermit
func List() ([]types.Container, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return list(client)
}

func list(client client.APIClient) ([]types.Container, error) {
	filters := filters.NewArgs()
	filters.Add(KermitLabel, "true")
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
		Filter: filters,
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}
