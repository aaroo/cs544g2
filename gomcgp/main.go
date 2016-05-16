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
			Name:  "hello",
			Usage: "get greeting",
			Action: func(c *cli.Context) {
				fmt.Printf("Hello\n")
			},
		},
		{
			Name:  "test",
			Usage: "test",
			Action: func(c *cli.Context) {
				fmt.Println("test 2")
			},
		},
	}
	app.Run(os.Args)
}
