package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/codegangsta/cli"
)

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
