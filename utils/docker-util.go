package utils

import (
	"fmt"
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

func (ctx *UtilContext) DeleteSingleContainer(container string, id bool, dry_run bool) error {

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
		return err //error("More then one option found with same id or name")
	}

	err = ctx.DeleteContainer(containers, "", dry_run)

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

	err = ctx.DeleteContainer(containers, "Exited", dry_run)

	return err
}

func (ctx *UtilContext) DeleteContainer(containerList []dockerclient.Container, status string, dry_run bool) (err error) {

	for _, c := range containerList {
		if strings.Contains(c.Status, status) {
			log.Println(c.Names)
			if !dry_run {
				err = ctx.client.RemoveContainer(c.Id, true, true)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	return err

}

func (ctx *UtilContext) RemoveDockerImage(image string, id bool) error {
	//FIXME: image delete for id

	imageInfo, err := ctx.client.RemoveImage(image)
	if err != nil {
		log.Printf("Unable to delete the Image: %v\n", err)
		return err
	}

	log.Println(imageInfo)
	return err

}

func (ctx *UtilContext) RemoveUntaggedDockerImages(dry_run bool) error {
	imageList, err := ctx.client.ListImages()
	if err != nil {
		log.Printf("Unable to list the images : %v\n", err)
		return err
	}

	for _, image := range imageList {
		if len(image.RepoTags) == 0 {
			log.Println(image.Id)
			// ctx.client.RemoveImage(image.Id)
		}
	}
	return err
}

func (ctx *UtilContext) CompactImage(id string) {
	log.Println("NOT IMPLEMENTED")
}
