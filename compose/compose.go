// Package compose aims to provide simple "helper" methods to ease the use of
// compose (through libcompose) in (integration) tests.
package compose

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/events"
	"github.com/docker/libcompose/project/options"
)

// Project holds compose related project attributes
type Project struct {
	composeProject project.APIProject
	name           string
	listenChan     chan events.Event
	started        chan struct{}
	stopped        chan struct{}
	deleted        chan struct{}
	client         client.APIClient
}

// CreateProject creates a compose project with the given name based on the
// specified compose files
func CreateProject(name string, composeFiles ...string) (*Project, error) {
	apiClient, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	composeProject, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: composeFiles,
			ProjectName:  name,
		},
	})
	if err != nil {
		return nil, err
	}
	p := &Project{
		composeProject: composeProject,
		name:           name,
		listenChan:     make(chan events.Event),
		started:        make(chan struct{}),
		stopped:        make(chan struct{}),
		deleted:        make(chan struct{}),
		client:         apiClient,
	}

	// Listen to compose events
	go p.startListening()
	p.composeProject.AddListener(p.listenChan)

	return p, nil
}

// Start creates and starts the compose project.
func (p *Project) Start() error {
	err := p.composeProject.Create(options.Create{})
	if err != nil {
		return err
	}
	err = p.composeProject.Start()
	if err != nil {
		return err
	}
	// Wait for compose to start
	<-p.started
	close(p.started)
	return nil
}

// Stop shuts down and clean the project
func (p *Project) Stop() error {
	// FIXME(vdemeester) Add options to customize the timeout
	err := p.composeProject.Stop(10)
	if err != nil {
		return err
	}
	<-p.stopped
	close(p.stopped)

	err = p.composeProject.Delete(options.Delete{})
	if err != nil {
		return err
	}
	<-p.deleted
	close(p.deleted)
	return nil
}

// Scale scale a service up
func (p *Project) Scale(service string, count int) error {
	return p.composeProject.Scale(10, map[string]int{
		service: count,
	})
}

func (p *Project) startListening() {
	for event := range p.listenChan {
		// FIXME Add a timeout on event ?
		if event.EventType == events.ProjectStartDone {
			p.started <- struct{}{}
		}
		if event.EventType == events.ProjectStopDone {
			p.stopped <- struct{}{}
		}
		if event.EventType == events.ProjectDeleteDone {
			p.deleted <- struct{}{}
		}
	}
}

// Containers lists containers for a given services.
func (p *Project) Containers(service string) ([]types.ContainerJSON, error) {
	containers := []types.ContainerJSON{}
	// Let's use engine-api for now as there is nothing really useful in
	// libcompose for now.
	filter := filters.NewArgs()
	filter.Add("label", "com.docker.compose.project="+p.name)
	filter.Add("label", "com.docker.compose.service="+service)
	containerList, err := p.client.ContainerList(context.Background(), types.ContainerListOptions{
		Filter: filter,
	})
	if err != nil {
		return containers, err
	}
	for _, c := range containerList {
		container, err := p.client.ContainerInspect(context.Background(), c.ID)
		if err != nil {
			return containers, err
		}
		containers = append(containers, container)
	}
	return containers, nil
}

// Container returns the one and only container for a given services. It returns an error
// if the service has more than one container (in case of scale)
func (p *Project) Container(service string) (types.ContainerJSON, error) {
	containers, err := p.Containers(service)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	if len(containers) > 1 {
		return types.ContainerJSON{}, fmt.Errorf("More than one container are running for '%s' service", service)
	}
	return containers[0], nil
}
