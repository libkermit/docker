package docker

import (
	"fmt"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
)

// Create lets you create a container with the specified image, and default
// configuration.
func Create(image string) (types.ContainerJSON, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return create(client, image, ContainerConfig{})
}

// CreateWithConfig lets you create a container with the specified image, and
// some custom simple configuration.
func CreateWithConfig(image string, config ContainerConfig) (types.ContainerJSON, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return create(client, image, config)
}

func create(client client.APIClient, image string, containerConfig ContainerConfig) (types.ContainerJSON, error) {
	config := &container.Config{
		Image:  image,
		Labels: KermitLabels,
	}

	containerName := fmt.Sprintf("kermit_%s", image)

	response, err := client.ContainerCreate(config, &container.HostConfig{}, nil, containerName)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return inspect(client, response.ID)
}
