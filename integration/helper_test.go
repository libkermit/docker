package dockerit

import (
	"testing"

	"golang.org/x/net/context"

	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	// "github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/filters"
	"github.com/libkermit/docker"
)

func setupTest(t *testing.T) *docker.Project {
	return cleanContainers(t)
}

func cleanContainers(t *testing.T) *docker.Project {
	client, err := dockerclient.NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}
	// FIXME(vdemeester) Fix this
	client.UpdateClientVersion(docker.CurrentAPIVersion)

	filterArgs := filters.NewArgs()
	if filterArgs, err = filters.ParseFlag(docker.KermitLabelFilter, filterArgs); err != nil {
		t.Fatal(err)
	}

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, container := range containers {
		t.Logf("cleaning container %s…", container.ID)
		if err := client.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
			Force: true,
		}); err != nil {
			t.Errorf("Error while removing container %s : %v\n", container.ID, err)
		}
	}

	return docker.NewProject(client)
}
