package composeit

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/go-check/check"
	compose "github.com/vdemeester/libkermit/compose/check"
	"github.com/vdemeester/libkermit/docker"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { check.TestingT(t) }

type CheckSuite struct{}

var _ = check.Suite(&CheckSuite{})

func (s *CheckSuite) TestSimpleProject(c *check.C) {
	project := compose.CreateProject(c, "simple", "../assets/simple.yml")
	project.Start(c)

	// FIXME(vdemeester) check that a container is running
	runningContainers, err := findContainersForProject("simple")
	c.Assert(err, check.IsNil)
	c.Assert(len(runningContainers), check.Equals, 1,
		check.Commentf("Expected 1 running container for this project, got %v", runningContainers))

	project.Stop(c)
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
