package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

func inspect(client client.APIClient, containerID string) (types.ContainerJSON, error) {
	container, err := client.ContainerInspect(containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return container, nil
}
