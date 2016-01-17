package docker

import (
	"testing"
)

func TestCreate(t *testing.T) {
	client := &fakeClient{}
	container, err := create(client, "notfound", ContainerConfig{})
	if err == nil {
		t.Fatalf("Expected an error, got nothing: %v", container)
	}

	client = &fakeClient{
		image:       "kermit",
		containerID: "containerID",
	}
	container, err = create(client, "kermit", ContainerConfig{})
	if err != nil {
		t.Fatalf("Didn't expect an error, got %v", err)
	}
	if container.ID != client.containerID {
		t.Fatalf("expected %s, got %s", client.containerID, container.ID)
	}
	if container.Name != "kermit_kermit" {
		t.Fatalf("expected %s, got %s", "kermit_kermit", container.Name)
	}
}
