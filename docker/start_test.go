package docker

import (
	"testing"
)

func TestStart(t *testing.T) {
	client := &fakeClient{}
	container, err := start(client, "notfound", ContainerConfig{})
	if err == nil {
		t.Fatalf("Expected an error, got nothing: %v", container)
	}

	client = &fakeClient{
		image:       "kermit",
		containerID: "containerID",
		running:     false,
	}
	container, err = start(client, "kermit", ContainerConfig{})
	if err == nil {
		t.Fatalf("Expected an error, got nothing: %v", container)
	}
	if container.ID != client.containerID {
		t.Fatalf("expected %s, got %s", client.containerID, container.ID)
	}
	if container.Name != "kermit_kermit" {
		t.Fatalf("expected %s, got %s", "kermit_kermit", container.Name)
	}
	if container.State.Running {
		t.Fatalf("expected the container to not be running, was not")
	}

	client = &fakeClient{
		image:       "kermit",
		containerID: "containerID",
		running:     true,
	}
	container, err = start(client, "kermit", ContainerConfig{})
	if err != nil {
		t.Fatalf("Didn't expect an error, got %v", err)
	}
	if container.ID != client.containerID {
		t.Fatalf("expected %s, got %s", client.containerID, container.ID)
	}
	if container.Name != "kermit_kermit" {
		t.Fatalf("expected %s, got %s", "kermit_kermit", container.Name)
	}
	if !container.State.Running {
		t.Fatalf("expected the container to be running, was not")
	}
}
