package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// Start lets you create and start a container with the specified image, and
// default configuration.
func Start(image string) (types.ContainerJSON, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return start(client, image, ContainerConfig{})
}

// StartWithConfig lets you create and start a container with the specified
// image, and some custom simple configuration.
func StartWithConfig(image string, config ContainerConfig) (types.ContainerJSON, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return start(client, image, config)
}

func start(client client.APIClient, image string, containerConfig ContainerConfig) (types.ContainerJSON, error) {
	container, err := create(client, image, containerConfig)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	if err := client.ContainerStart(container.ID); err != nil {
		return container, err
	}

	return inspect(client, container.ID)
}
