package docker

import (
	"golang.org/x/net/context"

	"github.com/docker/engine-api/types"
)

// Remove removes the container
func (p *Project) Remove(containerID string) error {
	if err := p.Client.ContainerRemove(context.Background(), types.ContainerRemoveOptions{
		ContainerID: containerID,
		Force:       true,
	}); err != nil {
		return err
	}
	return nil
}
