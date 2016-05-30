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
	app := cli.NewApp()
	app.Name = "gomcgp"
	app.Usage = "CLI client MCGP server"
	app.Version = buildver + " MCGP client (built " + buildtime + ")"
	app.Author = "Group 2"
	app.Email = "s@drexel.edu"
	app.EnableBashCompletion = true

	app.Before = func(c *cli.Context) error {
		// set safe random seed
		rand.Seed(time.Now().UTC().UnixNano())
		return nil
	}

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
		{
			Name:  "device",
			Usage: "device commands",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "retrieve list of devices",
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

						l_devices, err := clientDeviceList(conn)
						if err != nil {
							fmt.Printf("Error in list: %s\n", err)
							os.Exit(1)
						}
						fmt.Println("List successful")

						conn.Close()
						fmt.Println("Connection closed")

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
					Name:  "action",
					Usage: "perform an action on the device",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "open, on"},
						cli.BoolFlag{Name: "close, off"},
					},
					Action: func(c *cli.Context) error {
						if len(c.Args()) != 1 {
							fmt.Errorf("Provide device ID as first argument!\n")
							os.Exit(1)
						}
						if !c.Bool("open") && !c.Bool("close") {
							fmt.Errorf("Provide at least one action!\n")
							os.Exit(1)
						}

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

						var id, action int8
						if c.Bool("open") {
							action = STATUS_ON
						} else {
							action = STATUS_OFF
						}

						tmp, _ := strconv.ParseInt(c.Args().First(), 0, 8)
						id = int8(tmp)

						err := clientAction(conn, id, action)
						if err != nil {
							fmt.Printf("Error in action: %s\n", err)
							os.Exit(1)
						}
						fmt.Println("Action successful")

						conn.Close()
						fmt.Println("Connection closed")

						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
