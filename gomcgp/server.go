/**
CS 544 Computer Networks
6-1-2016

Group 2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou

Purpose:
	STATEFUL, CONCURRENT, SERVICE, UI
	server.go is the realization of the server end of the MCGP protocol.
	It will listen on designated port and handle the incoming connections
    in a stateful way. A new goroutine is spawned for each handled connection.
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	mrand "math/rand"
	"os"

	"github.com/codegangsta/cli"
)

// a helper function to check if error is not nil, print it and abort the server
func checkErrFatal(info string, err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s: %s\n", info, err)
		os.Exit(1)
	}
}

// function that starts the server listener
func runServer(c *cli.Context) error {
	// user-defined address:port to bind is passed as first argument
	bindTo := c.Args().First()
	// default address:port binding if unspecified
	if bindTo == "" {
		bindTo = "127.0.0.1:6666"
	}
	fmt.Printf("Binding server to %s\n", bindTo)

	ca := x509.NewCertPool()
	// load SSL CA certificate
	if pemData, err := ioutil.ReadFile(c.GlobalString("ca")); err != nil {
		log.Fatalf("Cannot load SSL CA certificate! %s", err)
	} else {
		ca.AppendCertsFromPEM(pemData)
	}

	// read and parse public/private key pair
	cert, err := tls.LoadX509KeyPair(c.GlobalString("certificate"), c.GlobalString("key"))
	checkErrFatal("Error loading certificate", err)

	// built TLS configuration
	config := tls.Config{
		Certificates:             []tls.Certificate{cert},
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                ca,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		}}

	// now start TLS listener (server socket)
	listener, err := tls.Listen("tcp", bindTo, &config)
	checkErrFatal("Error listening on "+bindTo, err)
	defer listener.Close() // defer closing the socket until end-of-function
	fmt.Println("Now listening...")
	// infinite loop to accept new connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		tlsconn, ok := conn.(*tls.Conn)
		tlsconn.Handshake() // TLS handshake
		if ok {
			fmt.Printf("Peer CN: %s\n", tlsconn.ConnectionState().PeerCertificates[0].Subject.CommonName)
		}

		// spawn handleClient function as a new go-routine for this connection
		go handleClient(tlsconn)
	}

	return nil
}

// a function that is run per-connection and executes the DFA
func handleClient(conn *tls.Conn) {
	state := "awaiting-version-handshake"
	var response PDU

	for {
		// get incoming PDU
		pdu, err := readPDU(conn)
		if err == io.EOF {
			conn.Close()
			fmt.Println("Connection closed.")
			return
		}
		// invalid PDU -> close connection
		if err != nil {
			fmt.Println("reading PDU failed:", err)
			conn.Close()
			return
		}

		// DFA states -> behavior dependant on current expectation (state)
		switch state {
		case "awaiting-version-handshake":
			{
				// if the version is supported, send back a supported_version response
				// and move to the next state
				if pdu.Version == SUPPORTED_VERSION && pdu.Operation == OP_VERSION {
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_VERSION}
					state = "awaiting-authentication"
				} else {
					// if the version is not supported, send back error
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_VERSION, Error: ERR_VERSION}
					// no change in state
				}
				if err := response.Write(conn); err != nil {
					fmt.Printf("Error sending response: %s\n", err)
					conn.Close()
					return
				}
			}
		case "awaiting-authentication":
			{
				fmt.Printf("Received ident (%d): %s\nmatch: %+v\n",
					len(pdu.Identifier), pdu.Identifier, pdu.Identifier == "john")

				// perform credential checks
				if pdu.Operation == OP_AUTHENTICATE && pdu.Identifier == "john" {
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_AUTHENTICATE}
					state = "idle"
				} else {
					// authentication error is sent
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_AUTHENTICATE, Error: ERR_AUTHENTICATE}
					// no change in state
				}
				if err := response.Write(conn); err != nil {
					fmt.Printf("Error sending response: %s\n", err)
					conn.Close()
					return
				}
			}
		case "idle":
			{
				if pdu.Operation == OP_LIST {
					// send device status list back to the client
					no_packets := int(math.Ceil(float64(len(devices)) / 5))
					for i := 0; i < no_packets; i++ {
						var l_devices [5]Device
						for j := 0; j < 5; j++ {
							if i*5+j < len(devices) {
								l_devices[j] = devices[i*5+j]
							}
						}
						response = PDU{Version: SUPPORTED_VERSION, Operation: OP_LIST, Devices: l_devices}
						if i+1 < no_packets {
							response.Operation = OP_LIST_CONTINUED
							if err := response.Write(conn); err != nil {
								fmt.Printf("Error sending response: %s\n", err)
								conn.Close()
								return
							}
						}
					}
				} else if pdu.Operation == OP_CONTROL {
					// perform the operation/action on the device as requested
					var success bool
					for k, v := range devices {
						if v.Id != pdu.Devices[0].Id {
							continue
						}

						if v.Type == TYPE_PRESSURE || v.Type == TYPE_TEMP {
							if pdu.Devices[0].Action == STATUS_ON {
								devices[k].Value = mrand.Float32() * 10
							} else {
								devices[k].Value = mrand.Float32() * 10
							}
						}
						devices[k].Status = pdu.Devices[0].Action
						success = true
					}

					// send back appropriate response (success/error)
					if success {
						response = PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL}
					} else {
						response = PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL, Error: ERR_CONTROL}
					}
				} else {
					// if it's neither OP_LIST nor OP_CONTROL, send back an error
					// saying the operation was unexpected as no other operations
					// can happen in the "idle" state

					response = PDU{Version: SUPPORTED_VERSION, Operation: pdu.Operation, Error: ERR_UNEXPECTED_OP}
				}
				if err := response.Write(conn); err != nil {
					fmt.Printf("Error sending response: %s\n", err)
					conn.Close()
					return
				}
			}
		}

	}
}
