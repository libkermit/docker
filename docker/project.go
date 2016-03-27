package docker

import (
	"github.com/docker/engine-api/client"
)

// Project holds docker related project attributes, like docker client, labels
// to put on the containers, and so on.
type Project struct {
	Client client.APIClient
	Labels map[string]string
}

// NewProjectFromEnv creates a project with a client that is build from environment variables.
func NewProjectFromEnv() (*Project, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return NewProject(client), nil
}

// NewProject creates a project with the given client and the default attributes.
func NewProject(client client.APIClient) *Project {
	return &Project{
		Client: client,
		Labels: KermitLabels,
	}
}
