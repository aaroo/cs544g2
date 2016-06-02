/**
CS 544 Computer Networks
6-1-2016
Group2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou
purpose:
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

//cleint connect using certficate
func clientConnect(c *cli.Context) (conn *tls.Conn) {
	ca := x509.NewCertPool()
	if pemData, err := ioutil.ReadFile(c.GlobalString("ca")); err != nil {
		log.Fatal("Cannot load SSL client CA certificate!")
	} else {
		ca.AppendCertsFromPEM(pemData)
	}
	cert, err := tls.LoadX509KeyPair(c.GlobalString("certificate"), c.GlobalString("key"))
	checkErrFatal("Error loading client certificate", err)
	config := tls.Config{RootCAs: ca, Certificates: []tls.Certificate{cert}}
	conn, err = tls.Dial("tcp", c.GlobalString("server"), &config)
	checkErrFatal("Error connecting to server", err)
	return
}
//handle client version handshake and returns errors as required
func clientVersionHandshake(conn *tls.Conn) (err error) {
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_VERSION}
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}
	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}

	fmt.Printf("Received handshake response: %+v\n", response)

	if response.Version != SUPPORTED_VERSION || response.Error != 0 {
		return errors.New("Unsuccessful Version Handshake")
	}
	return
}
//handles the client authentication
func clientAuthenticate(conn *tls.Conn, ident string) (err error) {
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_AUTHENTICATE, Identifier: ident}
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}
	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}

	fmt.Printf("Received authentication response: %+v\n", response)

	if response.Error != 0 {
		return errors.New("Unsuccessful Authentication")
	}
	return
}
//builds and recieves the list of devices sent from the server
func clientDeviceList(conn *tls.Conn) (l_devices []Device, err error) {
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_LIST}
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	response := PDU{Operation: OP_LIST_CONTINUED}
	for response.Operation == OP_LIST_CONTINUED {
		response, err = readPDU(conn)
		if err != nil {
			fmt.Println("reading response PDU failed:", err)
			return
		}

		fmt.Printf("Received list response: %+v\n", response)
		for _, v := range response.Devices {
			if v.Id != 0 {
				l_devices = append(l_devices, v)
			}
		}
	}

	if response.Error != 0 {
		return l_devices, errors.New("Unsuccessful List")
	}
	return
}
//client actions based on packet received from server
func clientAction(conn *tls.Conn, id, action int8) (err error) {
	var l_devices [5]Device
	l_devices[0].Id = id
	l_devices[0].Action = action
	vh := PDU{Version: SUPPORTED_VERSION, Operation: OP_CONTROL, Devices: l_devices}
	err = vh.Write(conn)
	if err != nil {
		fmt.Println("conn write failed:", err)
		return
	}

	response, err := readPDU(conn)
	if err != nil {
		fmt.Println("reading response PDU failed:", err)
		return
	}

	fmt.Printf("Received action response: %+v\n", response)

	if response.Error != 0 {
		return errors.New("Unsuccessful action")
	}
	return
}
