package composeit

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/libkermit/docker"
	compose "github.com/libkermit/docker/compose/testing"
)

func TestSimpleProject(t *testing.T) {
	project := compose.CreateProject(t, "simple", "../assets/simple.yml")
	project.Start(t)

	// FIXME(vdemeester) check that a container is running
	runningContainers, err := findContainersForProject("simple")
	if err != nil {
		t.Fatal(err)
	}
	if len(runningContainers) != 1 {
		t.Fatalf("Expected 1 running container for this project, got %v", runningContainers)
	}

	project.Stop(t)
}

func findContainersForProject(name string) ([]types.Container, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return []types.Container{}, err
	}
	filterArgs := filters.NewArgs()
	if filterArgs, err = filters.ParseFlag(docker.KermitLabelFilter, filterArgs); err != nil {
		return []types.Container{}, err
	}

	return client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
}
