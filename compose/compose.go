// Package compose aims to provide simple "helper" methods to ease the use of
// compose (through libcompose) in (integration) tests.
package compose

import (
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
)

// Project holds compose related project attributes
type Project struct {
	composeProject *project.Project
	listenChan     chan project.Event
	started        chan struct{}
	stopped        chan struct{}
	deleted        chan struct{}
}

// CreateProject creates a compose project with the given name based on the
// specified compose files
func CreateProject(name string, composeFiles ...string) (*Project, error) {
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
		listenChan:     make(chan project.Event),
		started:        make(chan struct{}),
		stopped:        make(chan struct{}),
		deleted:        make(chan struct{}),
	}

	// Listen to compose events
	go p.startListening()
	p.composeProject.AddListener(p.listenChan)

	return p, nil
}

// Start creates and starts the compose project.
func (p *Project) Start() error {
	err := p.composeProject.Create()
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
	err := p.composeProject.Down()
	if err != nil {
		return err
	}
	<-p.stopped
	close(p.stopped)

	err = p.composeProject.Delete()
	if err != nil {
		return err
	}
	<-p.deleted
	close(p.deleted)
	return nil
}

func (p *Project) startListening() {
	for event := range p.listenChan {
		// FIXME Add a timeout on event ?
		if event.EventType == project.EventProjectStartDone {
			p.started <- struct{}{}
		}
		if event.EventType == project.EventProjectDownDone {
			p.stopped <- struct{}{}
		}
		if event.EventType == project.EventProjectDeleteDone {
			p.deleted <- struct{}{}
		}
	}
}
