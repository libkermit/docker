package docker

import (
	"testing"
)

func TestStop(t *testing.T) {
	client := &fakeClient{}
	err := stopWithTimeout(client, "notfound", 0)
	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}

	client = &fakeClient{
		containerID: "containerID",
	}
	err = stopWithTimeout(client, "wrongID", 0)
	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}

	err = stopWithTimeout(client, "containerID", 0)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
}
