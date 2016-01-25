package docker

import (
	"testing"
)

func TestContainerStatus(t *testing.T) {
	client := &fakeClient{}
	_, err := containerStatus(client, "anyid", "running")
	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}
	client = &fakeClient{
		containerID: "containerID",
		state:       "paused",
	}
	_, err = containerStatus(client, "wrongID", "running")
	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}

	status, err := containerStatus(client, "containerID", "running")
	if err != nil {
		t.Fatalf("Didn't expect an error, got %v", err)
	}
	if status {
		t.Fatalf("expected %v, got %v", false, status)
	}

	status, err = containerStatus(client, "containerID", "paused")
	if err != nil {
		t.Fatalf("Didn't expect an error, got %v", err)
	}
	if !status {
		t.Fatalf("expected %v, got %v", true, status)
	}
}
