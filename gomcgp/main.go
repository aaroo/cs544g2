package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var (
	buildtime string
	buildver  string
)

func main() {

	var srcIPstr string = "127.0.0.1"

	app := cli.NewApp()
	app.Name = "gomcgp"
	app.Usage = "CLI client MCGP server"
	app.Version = buildver + " MCGP client (built " + buildtime + ")"
	app.Action = cli.ShowAppHelp
	app.Author = "Group 2"
	app.Email = ""
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config, c", Value: "mcgp.gcfg",
			Usage: "path to config file in gcfg format", EnvVar: "MCGP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "Test",
			Usage: "Test",
			Action: func(c *cli.Context) {
				fmt.Println("test 1")
			},
		},
		{
			Name:  "display",
			Usage: "Display Status of all devices",
			Action: func(c *cli.Context) {
				fmt.Println("test 2")
			},
		},
		{
			Name:  "ip",
			Usage: "Enter IP Address of Server",
			Action: func(c *cli.Context) error {
				srcIPstr = c.Args().First()
				fmt.Println("Server IP: ", srcIPstr)
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "options for server commands",
			Subcommands: []cli.Command{
				{
					Name:  "connect",
					Usage: "username",
					Action: func(c *cli.Context) error {
						fmt.Println("User: ", c.Args().First())
						fmt.Println("Server IP: ", srcIPstr)
						//send packet
						return nil
					},
				},
				{
					Name:  "disconnect",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("Disconnected: ", srcIPstr)
						//send disconnect
						return nil
					},
				},
			},
		},
		{
			Name:  "device",
			Usage: "options for device commans",
			Subcommands: []cli.Command{
				{
					Name:  "retrieve",
					Usage: "retrieve list of devices to file name",
					Action: func(c *cli.Context) error {
						fmt.Println("List of Devices stored to: ", c.Args().First())
						//send packet
						//wait for devices list
						return nil
					},
				},
				{
					Name:  "update",
					Usage: "updates device status",
					Action: func(c *cli.Context) error {
						fmt.Println("Update Device List")
						//send packet
						//wait for all devices
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
