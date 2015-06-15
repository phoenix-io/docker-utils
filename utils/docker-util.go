package utils

import (
	"github.com/samalba/dockerclient"
	"log"
	"fmt"
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

func (ctx *UtilContext) DeleteSingleContainer(container string, id bool) error {

	// Get the container
	var filter string
	if id {
		fmt.Sprintf(filter, "{'id':['%s']}", container)
	} else {
		fmt.Sprintf(filter, "{'name':[%s']}", container)
	}
	containers, err := ctx.client.ListContainers(true, false, filter)
	if err != nil {
		log.Fatal(err)
	}
	if dry_run {
		log.Println("dry-run - Skipping deletion of containers")
	}
	if len(containers) > 1 {
		log.Println("More then one option found with same id or name.")
		return error("More then one option found with same id or name")
	}

	err = DeleteContainer(containers, "")

	return err



}

func (ctx *UtilContext) DeleteExitedContainers(dry_run bool) error {

	// Get all containers
	containers, err := ctx.client.ListContainers(true, false, "")
	if err != nil {
		log.Fatal(err)
	}
	if dry_run {
		log.Println("dry-run - Skipping deletion of containers")
	}

	log.Println("Following containers will be deleted")

	err = DeleteContainer(containers, "Exited")

	return err
}

func (ctx *UtilContext) DeleteContainer(containerList *dockerclient.Container,status string) (err error) {

	for _, c := range containerList {
		if strings.Contains(c.Status, status) {
			log.Println(c.Names)
			if !dry_run {
				err = ctx.client.RemoveContainer(c.ID, true, true)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	return err

}

func (ctx *UtilContext) RemoveDockerImage(image string, id bool) error {

}

func (ctx *UtilContext) RemoveUntaggedDockerImages(dry_run bool) error {

}

func (ctx *UtilContext) CompactImage(id string) {

}
