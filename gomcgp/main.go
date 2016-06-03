/**
CS 544 Computer Networks
6-1-2016

Group 2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou

Purpose:
	UI
	main.go is the CLI "entrance" to the gomcgp program
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var (
	buildtime string
	buildver  string
)

func main() {
	// setup CLI
	app := cli.NewApp()
	app.Name = "gomcgp"
	app.Usage = "CLI client MCGP server"
	app.Version = "MCGP client " + buildver + " (built " + buildtime + ")"
	app.Author = "Group 2"
	app.Email = "s@drexel.edu"
	app.EnableBashCompletion = true

	app.Before = func(c *cli.Context) error {
		// set "safe" random seed based on current time
		rand.Seed(time.Now().UTC().UnixNano())
		return nil
	}

	// define CLI-global configuration flags
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "certificate, c", Value: "certificate.pem",
			Usage: "path to certificate in PEM format", EnvVar: "CERTIFICATE_PATH"},
		cli.StringFlag{Name: "key, k", Value: "key.pem",
			Usage: "path to key in PEM format", EnvVar: "KEY_PATH"},
		cli.StringFlag{Name: "ca", Value: "ca.pem",
			Usage: "path to CA in PEM format", EnvVar: "CA_PATH"},
		cli.StringFlag{Name: "server, s", Value: "localhost:6666",
			Usage: "server address in format host:port", EnvVar: "SERVER_ADDRESS"},
		cli.StringFlag{Name: "ident, i", Value: "john",
			Usage: "identity to use (CN of certificate)", EnvVar: "IDENT"},
	}

	// available commands in CLI
	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "s",
			Usage:     "start MGCP server",
			Action:    runServer,
		},
		/* dev-only
		{
			Name:  "debug",
			Usage: "Test connection",
			Action: func(c *cli.Context) error {
				conn := clientConnect(c)
				fmt.Println("Connection established")

				if err := clientVersionHandshake(conn); err != nil {
					fmt.Printf("Error in handshake: %s\n", err)
					os.Exit(1)
				}
				fmt.Println("Handshake successful")

				if err := clientAuthenticate(conn, c.GlobalString("ident")); err != nil {
					fmt.Printf("Error in authentication: %s\n", err)
					os.Exit(1)
				}
				fmt.Println("Authentication successful")

				conn.Close()
				fmt.Println("Connection closed")
				return nil
			},
		},
		*/
		{
			Name:      "device",
			ShortName: "d",
			Usage:     "device commands",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					ShortName: "l",
					Usage:     "retrieve list of devices",
					Action: func(c *cli.Context) error {
						// connect to the server
						conn := clientConnect(c)
						//fmt.Println("Connection established")

						// perform version handshake
						if err := clientVersionHandshake(conn); err != nil {
							fmt.Printf("Error in handshake: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("Handshake successful")

						// perform authentication
						if err := clientAuthenticate(conn, c.GlobalString("ident")); err != nil {
							fmt.Printf("Error in authentication: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("Authentication successful")

						// get devices
						l_devices, err := clientDeviceList(conn)
						if err != nil {
							fmt.Printf("Error in list: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("List successful")

						// close connection
						conn.Close()
						//fmt.Println("Connection closed")

						// draw the table with output
						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"ID", "Type", "Status", "Value"})
						for _, d := range l_devices {
							t := "unknown"
							switch d.Type {
							case TYPE_GARAGE_DOOR:
								{
									t = "garage door"
								}
							case TYPE_LIGHT:
								{
									t = "light"
								}
							case TYPE_PRESSURE:
								{
									t = "pressure"
								}
							case TYPE_TEMP:
								{
									t = "temperature"
								}
							}
							s := "unknown"
							switch d.Status {
							case STATUS_ON:
								{
									if d.Type == TYPE_GARAGE_DOOR {
										s = "open"
									} else {
										s = "on"
									}
								}
							case STATUS_OFF:
								{
									if d.Type == TYPE_GARAGE_DOOR {
										s = "closed"
									} else {
										s = "off"
									}
								}
							}

							table.Append([]string{
								fmt.Sprintf("%d", d.Id),
								t,
								s,
								fmt.Sprintf("%.4f", d.Value),
							})
						}
						table.Render()

						return nil
					},
				},
				{
					Name:      "action",
					ShortName: "a",
					Usage:     "perform an action on the device",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "open, on"},
						cli.BoolFlag{Name: "close, off"},
					},
					Action: func(c *cli.Context) error {
						// verify arguments
						if len(c.Args()) != 1 {
							fmt.Errorf("Provide device ID as first argument!\n")
							os.Exit(1)
						}

						// verify flags for sanity
						if !c.Bool("open") && !c.Bool("close") {
							fmt.Errorf("Provide at least one action!\n")
							os.Exit(1)
						}

						// connect to the server
						conn := clientConnect(c)
						//fmt.Println("Connection established")

						// perform version handshake
						if err := clientVersionHandshake(conn); err != nil {
							fmt.Printf("Error in handshake: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("Handshake successful")

						// perform authentication
						if err := clientAuthenticate(conn, c.GlobalString("ident")); err != nil {
							fmt.Printf("Error in authentication: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("Authentication successful")

						// prepare action parameter
						var id, action int8
						if c.Bool("open") {
							action = STATUS_ON
						} else {
							action = STATUS_OFF
						}

						// parse device ID from input
						tmp, _ := strconv.ParseInt(c.Args().First(), 0, 8)
						id = int8(tmp)

						// call the action on the server
						err := clientAction(conn, id, action)
						if err != nil {
							fmt.Printf("Error in action: %s\n", err)
							os.Exit(1)
						}
						//fmt.Println("Action successful")

						// close connection
						conn.Close()
						//fmt.Println("Connection closed")

						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
