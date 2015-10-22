package main

import (
	"github.com/phoenix-io/docker-utils/utils"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "docker-utils"
	app.Usage = "Toolchain for docker"
	app.Version = "0.2.0"
	app.Commands = []cli.Command{
		{
			Name:  "rmi",
			Usage: "deletes the docker images",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "dry",
					Usage: "[Optional] dry_run the command",
				},
				cli.BoolFlag{
					Name:  "untagged",
					Usage: "[Required] deletes untagged images",
				},
			},
			Action: func(c *cli.Context) {
				removeImages(c)
			},
		},
		{
			Name:  "rm",
			Usage: "deletes docker containers",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "exited",
					Usage: "[Required] deletes exited containers",
				},
				cli.BoolFlag{
					Name:  "dry",
					Usage: "[Optional] dry_run the command",
				},
				cli.IntFlag{
					Name:  "hours",
					Value: 24,
					Usage: "[Optional] hours before containers exited/stopped",
				},
			},
			Action: func(c *cli.Context) {
				removeContainers(c)
			},
		},
		{
			Name:  "flatten",
			Usage: "Compacts the images by flattening",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image",
					Usage: "[Required] Image file to flatten",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "[Required] New name of Image file",
				},
				cli.StringFlag{
					Name:  "tag",
					Usage: "[Required] tag for new image",
				},
			},
			Action: func(c *cli.Context) {
				flattenImage(c)
			},
		},
	}

	app.Run(os.Args)
}

// Intialze the docker client
func getUtilContext() *utils.UtilContext {
	ctx, err := utils.InitUtilContext()
	if err != nil {
		log.Println("Unable to initialze the dockerclient.")
		log.Println("Docker daemon should be running")
		return nil
	}
	return ctx
}

// remove untagged images
func removeImages(c *cli.Context) {
	dry := c.Bool("dry")
	untagged := c.Bool("untagged")

	if !untagged {
		cli.ShowCommandHelp(c, "rmi")
		fmt.Println("EXAMPLE:")
		fmt.Println("   command rmi --untagged")
		return
	}

	ctx := getUtilContext()
	if ctx == nil {
		return
	}

	ctx = getUtilContext()
	if ctx == nil {
		return
	}

	if untagged == true {
		ctx.RemoveUntaggedDockerImages(dry)
	}
	return
}

// delete containers which are not running
func removeContainers(c *cli.Context) {
	dry := c.Bool("dry")
	exited := c.Bool("exited")
	hours := c.Int("hours")

	if !exited {
		cli.ShowCommandHelp(c, "rm")
		fmt.Println("EXAMPLE:")
		fmt.Println("   command rm --exited")
		return
	}
	if hours <= 0 {
		hours = 24
	}

	ctx := getUtilContext()
	if ctx == nil {
		return
	}

	if exited == true {
		ctx.DeleteExitedContainers(dry, hours)
	}

	return
}

func flattenImage(c *cli.Context) {
	image := c.String("image")
	name := c.String("name")
	tag := c.String("tag")

	if image == "" && name == "" && tag == "" {
		cli.ShowCommandHelp(c, "flatten")
		fmt.Println("EXAMPLE:")
		fmt.Println("   command flatten --image debian --name debian-new --tag compact")
		return
	}
	ctx := getUtilContext()
	if ctx == nil {
		return
	}

	ctx.FlattenImage(image, name, tag)

	return
}
