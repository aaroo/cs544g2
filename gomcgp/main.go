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
	app := cli.NewApp()
	app.Name = "gomcgp"
	app.Usage = "CLI client MCGP server"
	app.Version = buildver + " MCGP client (built " + buildtime + ")"
	app.Author = "Group 2"
	app.Email = ""
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "certificate, c", Value: "certificate.pem",
			Usage: "path to certificate in PEM format", EnvVar: "CERTIFICATE_PATH"},
		cli.StringFlag{Name: "key, k", Value: "key.pem",
			Usage: "path to key in PEM format", EnvVar: "KEY_PATH"},
		cli.StringFlag{Name: "ca", Value: "ca.pem",
			Usage: "path to CA in PEM format", EnvVar: "CA_PATH"},
		cli.StringFlag{Name: "server, s", Value: "localhost:6666",
			Usage: "server address in format host:port", EnvVar: "SERVER_ADDRESS"},
	}

	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "s",
			Usage:     "start MGCP server",
			Action:    runServer,
		},
		{
			Name:  "debug",
			Usage: "Test connection",
			Action: func(c *cli.Context) error {
				conn := clientConnect(c)
				conn.Close()
				return nil
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
