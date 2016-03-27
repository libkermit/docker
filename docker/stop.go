package docker

import (
	"golang.org/x/net/context"
)

// Stop stops the container with a default timeout.
func (p *Project) Stop(containerID string) error {
	return p.StopWithTimeout(containerID, DefaultStopTimeout)
}

// StopWithTimeout stops the container with the specified timeout.
func (p *Project) StopWithTimeout(containerID string, timeout int) error {
	if err := p.Client.ContainerStop(context.Background(), containerID, timeout); err != nil {
		return err
	}

	return nil
}
