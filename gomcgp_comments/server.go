/**
CS 544 Computer Networks
6-1-2016
Group2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou
purpose:
	STATEFUL,CONCURRENT,SERVICE,UI
	server.go is the realization of the server end of the MCGP protocol.
	It will listen on designated port and handle the request it received
	from the client based on the DFA.
 */
package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	mrand "math/rand"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

/**
 Function to check fatal error, if error happened, stop the server.
 */
func checkErrFatal(info string, err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s: %s\n", info, err)
		os.Exit(1)
	}
}
/**
Function to start and configure the server.
 */
func runServer(c *cli.Context) error {
	//bind server to the specific ip address and port.
	bindTo := c.Args().First()
	//if user did not designate ip address and port, bind server to the default ip address.
	if bindTo == "" {
		bindTo = "127.0.0.1:6666"
	}
	fmt.Printf("Binding server to %s\n", bindTo)

	//generate a new, empty CertPool.
	ca := x509.NewCertPool()

	//load SSL client's CA certificate.
	if pemData, err := ioutil.ReadFile(c.GlobalString("ca")); err != nil {
		log.Fatalf("Cannot load SSL client CA certificate! %s", err)
	} else {
		ca.AppendCertsFromPEM(pemData)
	}

	//read and parse public/private key pair
	cert, err := tls.LoadX509KeyPair(c.GlobalString("certificate"), c.GlobalString("key"))
	checkErrFatal("Error loading certificate", err)
	//if no error happened, set up the tls configuration.
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

	//set up the time.
	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	//set up tls listener, the listener will listen on specific port and ip address
	//that the server has been bound to.
	listener, err := tls.Listen("tcp", bindTo, &config)
	checkErrFatal("Error listening on "+bindTo, err)
	defer listener.Close()
	fmt.Println("Now listening...")
	//if no error happened, the listener will keep on listening for all coming connections.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		tlsconn, ok := conn.(*tls.Conn)
		tlsconn.Handshake()
		if ok {
			//fmt.Printf("%+v\n", tlsconn.ConnectionState().PeerCertificates[0])
			fmt.Printf("Peer CN: %s\n", tlsconn.ConnectionState().PeerCertificates[0].Subject.CommonName)
		}
		go handleClient(tlsconn)
	}

	return nil
}


/**
Function that will handle the client's connection.
 */
func handleClient(conn *tls.Conn) {
	state := "awaiting-version-handshake"
	var response PDU

	for {
		//parse the PDU sent by the client.
		pdu, err := readPDU(conn)
		if err == io.EOF {
			conn.Close()
			fmt.Println("Connection closed.")
			return
		}
		if err != nil {
			fmt.Println("reading PDU failed:", err)
			conn.Close()
			return
		}

		//DFA that the server will follow to deal with the parsed pdu.
		switch state {
		case "awaiting-version-handshake":
			{
				//if the version is supported, send back a supported_version response
				//and move to the next state.
				if pdu.Version == SUPPORTED_VERSION && pdu.Operation == OP_VERSION {
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_VERSION}
					state = "awaiting-authentication"
				} else {
					//if the version is not supported, send back a error_version response and close the connection.
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
				// if identifiers mateches, move to the next states.
				if pdu.Operation == OP_AUTHENTICATE && pdu.Identifier == "john" {
					response = PDU{Version: SUPPORTED_VERSION, Operation: OP_AUTHENTICATE}
					state = "idle"
				} else {
					//if identifiers do not match, send back a authentication error and close the connection.
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
				// send device status list back to the client
				if pdu.Operation == OP_LIST {
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
				}else if pdu.Operation == OP_CONTROL {
					//change device's status based on the received PDU.
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
					if success {
						response = PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL}
					} else {
						response = PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL, Error: ERR_CONTROL}
					}
				} else {
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
