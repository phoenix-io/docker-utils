package utils

import (
	"github.com/samalba/dockerclient"
	"log"
	"strings"
)

func RemoveDockerContainers(test bool) error {
	// Init the client
	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get only running containers
	containers, err := docker.ListContainers(true, false, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Following containers will be deleted")
	for _, c := range containers {
		if strings.Contains(c.Status, "Exit") {
			log.Println(c.Names)
		}
	}
	if test {
		log.Println("Test - No action taken")
	}
	return err
}
