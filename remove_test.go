package docker

import (
	"testing"

	"github.com/libkermit/docker/test"
)

type RemoveClient struct {
	test.NopClient
}

func TestProjectRemoteError(t *testing.T) {
	project := NewProject(&RemoveClient{})
	err := project.Remove("my_container")
	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}
}
