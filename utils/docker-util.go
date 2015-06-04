package utils

import (
	"github.com/samalba/dockerclient"
	"log"
	"strings"
)

type UtilContext struct {
	client *dockerclient.DockerClient
}

func InitUtilContext() (*UtilContext, error) {
	client, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	context := &UtilContext{client: client}
	return context, nil

}

func (ctx *UtilContext) RemoveDockerContainers(test bool) error {

	// Get all containers
	containers, err := ctx.client.ListContainers(true, false, "")
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

func (ctx *UtilContext) RemoveDockerImages(test bool) error {

}

func (ctx *UtilContext) CompactImage(id string) {

}
