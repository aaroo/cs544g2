/**
CS 544 Computer Networks
6-1-2016

Group 2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou

Purpose:
	CLIENT, STATEFUL, UI
	client.go is the realization of the server end of the MCGP protocol.
	It will send on designated port and handle the data it received
	from the server based on the DFA.
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/codegangsta/cli"
)

// This function opens a new TCP+TLS connection and return its "handle".
// It will perform TLS client authentication using the given certificate.
// It pulls all the configuration from the CLI context it's given.
// Fails fatally because it's client-side.
func clientConnect(c *cli.Context) (conn *tls.Conn) {
	ca := x509.NewCertPool()
	// Need to load the CA certificate to be able to establish authenticity
	// of the server.
	if pemData, err := ioutil.ReadFile(c.GlobalString("ca")); err != nil {
		log.Fatal("Cannot load SSL client CA certificate!")
	} else {
		ca.AppendCertsFromPEM(pemData)
	}

	// Now loading the client's certificate
	cert, err := tls.LoadX509KeyPair(c.GlobalString("certificate"), c.GlobalString("key"))
	checkErrFatal("Error loading client certificate", err)
	config := tls.Config{RootCAs: ca, Certificates: []tls.Certificate{cert}}

	// Dialing (connecting)...
	conn, err = tls.Dial("tcp", c.GlobalString("server"), &config)
	checkErrFatal("Error connecting to server", err)
	return
}

// Move through the version-handshake step
func clientVersionHandshake(conn *tls.Conn) (err error) {
	// build a packet
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_VERSION}

	// send a packet through the socket
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	// get the response
	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}
	//fmt.Printf("Received handshake response: %+v\n", response)

	// verify the response
	if response.Version != SUPPORTED_VERSION || response.Error != 0 {
		return errors.New("Unsuccessful Version Handshake")
	}
	return
}

// Move through the authentication step
func clientAuthenticate(conn *tls.Conn, ident string) (err error) {
	// build a packet
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_AUTHENTICATE, Identifier: ident}

	// send a packet through the socket
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	// get the response
	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}
	//fmt.Printf("Received authentication response: %+v\n", response)

	// verify the response
	if response.Error != 0 {
		return errors.New("Unsuccessful Authentication")
	}
	return
}

// Request and receive a full list of devices from the server.
// May receive and bundle-up multiple packets from the server and combine
// them into one logical list.
func clientDeviceList(conn *tls.Conn) (l_devices []Device, err error) {
	// build a packet
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_LIST}

	// send a packet through the socket
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	// get the response
	response := PDU{Operation: OP_LIST_CONTINUED}

	// check if there is more packets following
	for response.Operation == OP_LIST_CONTINUED {
		// read another packet
		response, err = readPDU(conn)
		if err != nil {
			fmt.Println("reading response PDU failed:", err)
			return
		}
		//fmt.Printf("Received list response: %+v\n", response)

		// append to local array
		for _, v := range response.Devices {
			if v.Id != 0 {
				l_devices = append(l_devices, v)
			}
		}
	}

	// verify the response
	if response.Error != 0 {
		return l_devices, errors.New("Unsuccessful List")
	}
	return
}

// Perform an action on a device on the server
func clientAction(conn *tls.Conn, id, action int8) (err error) {
	// build a packet
	var l_devices [5]Device
	l_devices[0].Id = id
	l_devices[0].Action = action
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL, Devices: l_devices}

	// send a packet through the socket
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	// get the response
	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}
	//fmt.Printf("Received action response: %+v\n", response)

	// verify the response
	if response.Error != 0 {
		return errors.New("Unsuccessful action")
	}
	return
}
