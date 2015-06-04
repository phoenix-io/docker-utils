package main

import (
	"github.com/UmbrellaOps/docker-utils/utils"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "docker-utils"
	app.Usage = "Toolchain for docker"
	app.Commands = []cli.Command{
		{
			Name:  "clean",
			Usage: "deletes the docker images",
			Subcommands: []cli.Command{
				{
					Name:    "images",
					Aliases: []string{"I"},
					Usage:   "remove untag Docker Images",
					Action: func(c *cli.Context) {
						removeImages(c)
					},
				},
				{
					Name:    "container",
					Aliases: []string{"C"},
					Usage:   "remove an stopped containers",
					Action: func(c *cli.Context) {
						removeContainers(c)
					},
				},
			},
		},
		{
			Name:  "compactimages",
			Usage: "Compacts the images by flaterning",
			Action: func(c *cli.Context) {
				log.Println("Compacts the images by flterning")
			},
		},
	}

	app.Run(os.Args)
}

// remove images based on given arguments
func removeImages(c *cli.Context) bool {
	log.Println("NOT IMPLEMENTED")
	return false
}

// delete containers which are not running
func removeContainers(c *cli.Context) bool {
	utils.RemoveDockerContainers(true)
	return false

}
