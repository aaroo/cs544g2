package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func checkErrFatal(info string, err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s: %s\n", info, err)
		os.Exit(1)
	}
}

func runServer(c *cli.Context) error {
	bindTo := c.Args().First()
	if bindTo == "" {
		bindTo = "127.0.0.1:6666"
	}
	fmt.Printf("Binding server to %s\n", bindTo)

	ca := x509.NewCertPool()

	if pemData, err := ioutil.ReadFile(c.GlobalString("ca")); err != nil {
		log.Fatalf("Cannot load SSL client CA certificate! %s", err)
	} else {
		ca.AppendCertsFromPEM(pemData)
	}

	cert, err := tls.LoadX509KeyPair(c.GlobalString("certificate"), c.GlobalString("key"))
	checkErrFatal("Error loading certificate", err)
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

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", bindTo, &config)
	checkErrFatal("Error listening on "+bindTo, err)
	fmt.Println("Now listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		defer conn.Close()
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		tlsconn, ok := conn.(*tls.Conn)
		tlsconn.Handshake()
		if ok {
			fmt.Printf("Peer CN: %s\n", tlsconn.ConnectionState().PeerCertificates[0].Subject.CommonName)
		}
		go handleClient(conn)
	}

	return nil
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
