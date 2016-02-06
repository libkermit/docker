package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// Inspect returns the container informations
func Inspect(containerID string) (types.ContainerJSON, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return inspect(client, containerID)
}

func inspect(client client.APIClient, containerID string) (types.ContainerJSON, error) {
	container, err := client.ContainerInspect(containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return container, nil
}
