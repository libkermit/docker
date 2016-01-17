package docker

import (
	"github.com/docker/engine-api/client"
)

// Stop stops the container with a default timeout.
func Stop(containerID string) error {
	return StopWithTimeout(containerID, DefaultStopTimeout)
}

// StopWithTimeout stops the container with the specified timeout.
func StopWithTimeout(containerID string, timeout int) error {
	client, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	return stopWithTimeout(client, containerID, timeout)
}

func stopWithTimeout(client client.APIClient, containerID string, timeout int) error {
	if err := client.ContainerStop(containerID, timeout); err != nil {
		return err
	}

	return nil
}
