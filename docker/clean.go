package docker

import (
	"github.com/docker/engine-api/client"
)

// Clean stops and removes (by default, controllable with the keep) kermit containers
func Clean(keep bool) error {
	client, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	return clean(client, keep)
}

func clean(client client.APIClient, keep bool) error {
	containers, err := list(client)
	if err != nil {
		return err
	}
	for _, container := range containers {
		if err := stopWithTimeout(client, container.ID, 1); err != nil {
			return err
		}
		if !keep {
			if err := remove(client, container.ID); err != nil {
				return err
			}
		}
	}
	return nil
}
