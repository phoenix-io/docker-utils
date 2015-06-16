package main

import (
	_ "github.com/UmbrellaOps/docker-utils/utils"
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
					Aliases: []string{"i"},
					Usage:   "remove untag Docker Images",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "dry",
							Usage: "dry_run the command",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "dry_run the command",
						},
						cli.StringFlag{
							Name:  "id",
							Usage: "dry_run the command",
						},
					},
					Action: func(c *cli.Context) {
						removeImages(c)
					},
				},
				{
					Name:    "container",
					Aliases: []string{"c"},
					Usage:   "remove an stopped containers",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "dry",
							Usage: "dry_run the command",
						},
					},
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
	log.Println(c.String("dry"))
	log.Println(c.String("name"))
	log.Println(c.String("id"))
	log.Println("NOT IMPLEMENTED")
	return false
}

// delete containers which are not running
func removeContainers(c *cli.Context) bool {
	//	utils.DeleteSingleContainer("aaa",true,true)
	return false

}
