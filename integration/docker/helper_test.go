package dockerit

import (
	"testing"

	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	// "github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/filters"
	"github.com/vdemeester/libkermit/docker"
)

func setupTest(t *testing.T) {
	cleanContainers(t)
}

func cleanContainers(t *testing.T) {
	client, err := dockerclient.NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}

	filterArgs := filters.NewArgs()
	if filterArgs, err = filters.ParseFlag(docker.KermitLabelFilter, filterArgs); err != nil {
		t.Fatal(err)
	}

	containers, err := client.ContainerList(types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, container := range containers {
		t.Logf("cleaning container %s…", container.ID)
		if err := client.ContainerRemove(types.ContainerRemoveOptions{
			ContainerID: container.ID,
			Force:       true,
		}); err != nil {
			t.Errorf("Error while removing container %s : %v\n", container.ID, err)
		}
	}
}
