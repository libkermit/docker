package dockerit

import (
	"testing"

	"github.com/vdemeester/libkermit/docker"
)

func TestCreateSimple(t *testing.T) {
	setupTest(t)

	container, err := docker.Create("busybox")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if container.ID == "" {
		t.Fatalf("expected a containerId, got nothing")
	}
}

func TestStartAndStop(t *testing.T) {
	setupTest(t)

	container, err := docker.Start("busybox")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if container.ID == "" {
		t.Fatalf("expected a containerId, got nothing")
	}

	err = docker.Stop(container.ID)
	if err != nil {
		t.Fatal(err)
	}

}
