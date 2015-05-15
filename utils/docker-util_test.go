package utils

import "testing"

func TestRemoveDockerContainers(t *testing.T) {
	err := RemoveDockerContainers(true)
	if err != nil {
		t.Errorf("Failed to remove Docker Containers %v", err)
	}

}
