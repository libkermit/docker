package docker

import (
	"testing"
)

func TestInspect(t *testing.T) {
	client := &fakeClient{}
	container, err := inspect(client, "anyid")
	if err == nil {
		t.Fatalf("Expected an error, got %v", container)
	}

	client = &fakeClient{
		containerID: "containerID",
	}
	container, err = inspect(client, "wrongID")
	if err == nil {
		t.Fatalf("Expected an error, got %v", container)
	}

	container, err = inspect(client, "containerID")
	if err != nil {
		t.Fatalf("Didn't expect an error, got %v", err)
	}
	if container.ID != client.containerID {
		t.Fatalf("expected %s, got %s", client.containerID, container.ID)
	}
}
