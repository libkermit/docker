package dockerit

import (
	"golang.org/x/net/context"

	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/go-check/check"
	d "github.com/vdemeester/libkermit/docker"
	docker "github.com/vdemeester/libkermit/docker/check"
)

func setupTest(c *check.C) *docker.Project {
	return cleanContainers(c)
}

func cleanContainers(c *check.C) *docker.Project {
	client, err := dockerclient.NewEnvClient()
	c.Assert(err, check.IsNil)

	filterArgs := filters.NewArgs()
	filterArgs, err = filters.ParseFlag(d.KermitLabelFilter, filterArgs)
	c.Assert(err, check.IsNil)

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
	c.Assert(err, check.IsNil)

	for _, container := range containers {
		c.Logf("cleaning container %s…", container.ID)
		if err := client.ContainerRemove(context.Background(), types.ContainerRemoveOptions{
			ContainerID: container.ID,
			Force:       true,
		}); err != nil {
			c.Errorf("Error while removing container %s : %v\n", container.ID, err)
		}
	}

	return docker.NewProject(client)
}
