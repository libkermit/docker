package composeit

import (
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/vdemeester/libkermit/compose"
	"github.com/vdemeester/libkermit/docker"
)

func TestSimpleProject(t *testing.T) {
	project, err := compose.CreateProject("simple", "./assets/simple.yml")
	if err != nil {
		t.Fatal(err)
	}

	err = project.Start()
	if err != nil {
		t.Fatal(err)
	}

	// FIXME(vdemeester) check that a container is running
	runningContainers, err := findContainersForProject("simple")
	if err != nil {
		t.Fatal(err)
	}
	if len(runningContainers) != 1 {
		t.Fatalf("Expected 1 running container for this project, got %v", runningContainers)
	}

	err = project.Stop()
	if err != nil {
		t.Fatal(err)
	}
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

	return client.ContainerList(types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
}
