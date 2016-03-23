package docker

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/strslice"
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
	err := ensureImageExists(client, image)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	labels := mergeLabels(KermitLabels, containerConfig.Labels)
	config := &container.Config{
		Image:  image,
		Labels: labels,
	}

	if len(containerConfig.Entrypoint) > 0 {
		config.Entrypoint = strslice.StrSlice(containerConfig.Entrypoint)
	}
	if len(containerConfig.Cmd) > 0 {
		config.Cmd = strslice.StrSlice(containerConfig.Cmd)
	}

	var containerName string
	if containerConfig.Name != "" {
		containerName = containerConfig.Name
	} else {
		containerName = fmt.Sprintf("kermit_%s", namesgenerator.GetRandomName(10))
	}

	response, err := client.ContainerCreate(context.Background(), config, &container.HostConfig{}, nil, containerName)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return inspect(client, response.ID)
}

func mergeLabels(defaultLabels, additionnalLabels map[string]string) map[string]string {
	labels := make(map[string]string, len(defaultLabels))
	if len(additionnalLabels) > 0 {
		for key, value := range additionnalLabels {
			labels[key] = value
		}
	}
	// default labels overrides additionnals
	for key, value := range defaultLabels {
		labels[key] = value
	}
	return labels
}
