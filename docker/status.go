package docker

import (
	"github.com/docker/engine-api/client"
)

// IsRunning checks if the container is running or not
func IsRunning(containerID string) (bool, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return false, nil
	}
	return containerStatus(client, containerID, "running")
}

// IsStopped checks if the container is running or not
func IsStopped(containerID string) (bool, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return false, nil
	}
	return containerStatus(client, containerID, "stopped")
}

// IsPaused checks if the container is running or not
func IsPaused(containerID string) (bool, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return false, nil
	}
	return containerStatus(client, containerID, "paused")
}

func containerStatus(client client.APIClient, containerID, status string) (bool, error) {
	containerJSON, err := inspect(client, containerID)
	if err != nil {
		return false, err
	}
	return containerJSON.State.Status == status, nil

}
