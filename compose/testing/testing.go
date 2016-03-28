// Package testing aims to provide simple "helper" methods to ease the use of
// compose (through libcompose) in (integration) tests using built-in testing.
package testing

import (
	"testing"

	"github.com/vdemeester/libkermit/compose"
)

// Project holds compose related project attributes
type Project struct {
	project *compose.Project
}

// CreateProject creates a compose project with the given name based on the
// specified compose files
func CreateProject(t *testing.T, name string, composeFiles ...string) *Project {
	project, err := compose.CreateProject(name, composeFiles...)
	if err != nil {
		t.Fatalf("error while creating compose project %s from %v: %s", name, composeFiles, err.Error())
	}
	return &Project{
		project: project,
	}
}

// Start creates and starts the compose project.
func (p *Project) Start(t *testing.T) {
	if err := p.project.Start(); err != nil {
		t.Fatalf("error while starting compose project: %s", err.Error())
	}
}

// Stop shuts down and clean the project
func (p *Project) Stop(t *testing.T) {
	if err := p.project.Stop(); err != nil {
		t.Fatalf("error while stopping compose project: %s", err.Error())
	}
}
