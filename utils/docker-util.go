package utils

import (
	"errors"
	"fmt"
	"github.com/samalba/dockerclient"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type UtilContext struct {
	client *dockerclient.DockerClient
}

func InitUtilContext() (*UtilContext, error) {
	client, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	context := &UtilContext{client: client}
	return context, nil

}

func (ctx *UtilContext) DeleteExitedContainers(dry_run bool) error {

	// Get all containers
	containers, err := ctx.client.ListContainers(true, false, "")
	if err != nil {
		fmt.Println("Error while retriving conatiner list")
		return err
	}

	err = ctx.DeleteContainer(containers, "Exited", dry_run)
	if dry_run {
		fmt.Println("dry-run - Skipped deletion of containers")
	}

	return nil
}

func (ctx *UtilContext) DeleteContainer(containerList []dockerclient.Container, status string, dry_run bool) (err error) {

	count := 0

	for _, c := range containerList {
		if !strings.Contains(c.Status, status) {
			continue
		}

		count = count + 1
		fmt.Printf("Deleting container : %s\n", c.Names)
		if dry_run {
			continue
		}

		err = ctx.client.RemoveContainer(c.Id, true, true)
		printError("Error while deleting container", err)
	}

	if count == 0 {
		fmt.Printf("\nNo container elegible for deletion\n\n")
	}
	return err
}

func (ctx *UtilContext) RemoveUntaggedDockerImages(dry_run bool) error {

	count := 0

	imageList, err := ctx.client.ListImages()
	if err != nil {
		fmt.Printf("Unable to featch list of images\n")
		return err
	}

	for _, image := range imageList {
		// Check for untagged images.
		if !strings.EqualFold(image.RepoTags[0], "<none>:<none>") {
			continue
		}

		count = count + 1
		fmt.Printf("Removing image : %s \n", image.Id)
		if dry_run {
			continue
		}
		_, err = ctx.client.RemoveImage(image.Id)
		printError("Unable to delete Image\n", err)
	}

	if count == 0 {
		fmt.Printf("\nNo docker image elegible for deletion\n\n\n")
	}
	if dry_run {
		fmt.Println("dry-run - Skipped deletion of Images")
	}
	return err
}

func (ctx *UtilContext) FlattenImage(image string, name string, tag string) error {

	installed := checkInstalled("tar")
	if !installed {
		fmt.Printf("tar is not installed on system \ntar is required for this feature\n")
		return errors.New("Tar not found!")
	}
	//Create a temp folder.
	dir, err := ioutil.TempDir("/tmp", normalizeStr(image))
	if err != nil {
		fmt.Println("Error: Unable to create temp folder.")
		return err
	}
	defer os.RemoveAll(dir)

	fmt.Println("This may take few minutes, depending on size image")
	fmt.Println("Exporting image ....")
	// Export image in temp folder
	tarfile, err := ctx.exportImage(image, dir)
	if err != nil {
		fmt.Printf("Error while exporting Image \n")
		return err
	}
	fmt.Printf("Export completed successfully!\n")
	fmt.Println("Importing image in flatten mode.....")
	// Import tar file to docker repo
	err = ctx.importImage(tarfile, name, tag)
	if err != nil {
		fmt.Println("Unable to import, exported image")
		return err
	}
	fmt.Printf("Image imported sucessfully as %s:%s\n\n", name, tag)
	fmt.Printf("Image flattening completed. \nPlease run \"docker images\" to verify")

	return nil
}

func (ctx *UtilContext) exportImage(image string, dir string) (string, error) {
	// Create a instance
	// store id
	// on exit of container, do export.
	config := dockerclient.ContainerConfig{Image: image, Cmd: []string{"/bin/bash"}}
	id, err := ctx.client.CreateContainer(&config, "")
	if err != nil {
		fmt.Printf("Error while create container \n")
		return "", err
	}
	cmd := exec.Command("docker", "export", id)
	// Temp output file.
	filepath := dir + "/image.tar"
	// open the file for writing
	tarfile, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Unable to create temp file\n")
		return "", err
	}
	defer tarfile.Close()

	//Pipe the output of "docker export <id> to tarfile.
	cmd.Stdout = tarfile

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Unable to export image\n")
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Unable to export image\n")
		return "", err
	}

	return filepath, err
}

func (ctx *UtilContext) importImage(filepath string, name string, tag string) error {
	// Open file with io.Reader
	// Import the file.
	tarfile, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Unable to open file for import \n")
		return err
	}
	defer tarfile.Close()

	_, err = ctx.client.ImportImage("", name, tag, tarfile)
	if err != nil {
		fmt.Printf("Unable to import image\n")
		return err
	}

	return nil
}

// ------------------- Helper functions -------------------

// returns a name that doesnt have invalid seperator
func normalizeStr(str string) string {
	return strings.Replace(str, "/", "-", -1)
}

func printError(msg string, err error) bool {
	if err != nil {
		fmt.Println(msg)
		return true
	}
	return false
}

func checkInstalled(pkg string) bool {
	_, err := exec.LookPath(pkg)
	if err != nil {
		return false
	}
	return true
}
