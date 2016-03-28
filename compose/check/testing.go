// Package testing aims to provide simple "helper" methods to ease the use of
// compose (through libcompose) in (integration) tests using the go-check package.
package check

import (
	"github.com/go-check/check"

	"github.com/vdemeester/libkermit/compose"
)

// Project holds compose related project attributes
type Project struct {
	project *compose.Project
}

// CreateProject creates a compose project with the given name based on the
// specified compose files
func CreateProject(c *check.C, name string, composeFiles ...string) *Project {
	project, err := compose.CreateProject(name, composeFiles...)
	c.Assert(err, check.IsNil,
		check.Commentf("error while creating compose project %s from %v", name, composeFiles))
	return &Project{
		project: project,
	}
}

// Start creates and starts the compose project.
func (p *Project) Start(c *check.C) {
	c.Assert(p.project.Start(), check.IsNil,
		check.Commentf("error while starting compose project"))
}

// Stop shuts down and clean the project
func (p *Project) Stop(c *check.C) {
	c.Assert(p.project.Stop(), check.IsNil,
		check.Commentf("error while stopping compose project"))
}
