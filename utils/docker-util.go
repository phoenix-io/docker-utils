package utils

import "testing"

func TestInitUtilContext(t *testing.T) {
	_, err := InitUtilContext()
	if err != nil {
		t.Errorf("Failed to get util-context %v", err)
	}
}

func TestRemoveUntaggedDockerImages(t *testing.T) {
	ctx, _ := InitUtilContext()
	err := ctx.RemoveUntaggedDockerImages(true)
	if err != nil {
		t.Errorf("Failed to remove Docker images %v", err)
	}
}

func TestDeleteExitedContainers(t *testing.T) {
	ctx, _ := InitUtilContext()
	err := ctx.DeleteExitedContainers(true)
	if err != nil {
		t.Errorf("Failed to delete exited Containers")
	}
}

func TestFlattenImage(t *testing.T) {
	ctx, _ := InitUtilContext()
	err := ctx.FlattenImage("debian:latest", "myubuntu", "temp")
	if err != nil {
		t.Errorf("Failed to flatten the image")
	}

}
