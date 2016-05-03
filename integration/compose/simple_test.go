package composeit

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/libkermit/docker"
	"github.com/libkermit/docker/compose"
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

func TestProjectContainers(t *testing.T) {
	project, err := compose.CreateProject("simple", "./assets/simple.yml")
	if err != nil {
		t.Fatal(err)
	}

	if err := project.Start(); err != nil {
		t.Fatal(err)
	}

	if err := project.Scale("hello", 2); err != nil {
		t.Fatal(err)
	}

	containers, err := project.Containers("hello")
	if err != nil {
		t.Fatal(err)
	}

	if len(containers) != 2 {
		t.Fatalf("expected 2 containers, got %v", containers)
	}

	if err := project.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestProjectContainer(t *testing.T) {
	project, err := compose.CreateProject("simple", "./assets/simple.yml")
	if err != nil {
		t.Fatal(err)
	}

	if err := project.Start(); err != nil {
		t.Fatal(err)
	}

	container, err := project.Container("hello")
	if err != nil {
		t.Fatal(err)
	}
	if container.Name != "/simple_hello_1" {
		t.Fatalf("expected name '/simple_hello_1', got %s", container.Name)
	}

	if err := project.Scale("hello", 2); err != nil {
		t.Fatal(err)
	}

	_, err = project.Container("hello")
	if err == nil {
		t.Fatalf("expected an error on getting one container, got nothing")
	}

	if err := project.Stop(); err != nil {
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

	return client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
}
