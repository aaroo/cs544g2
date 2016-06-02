/**
CS 544 Computer Networks
6-1-2016
Group2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou
purpose:
	SERVICE 
	protocol.go is the realization of the MCGP protocol.
	It will send on build and send the packets per the DFA
 */
 
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)
//build packet information and states per DFA
const (
	SUPPORTED_VERSION = 1
)
const (
	_                 = iota
	OP_VERSION        // 0x01
	OP_AUTHENTICATE   // 0x02
	OP_LIST           // 0x03
	OP_LIST_CONTINUED // 0x04
	OP_CONTROL        // 0x05
)
const (
	_                 = iota
	ERR_VERSION       // 0x01
	ERR_AUTHENTICATE  // 0x02
	ERR_LIST          // 0x03
	ERR_CONTROL       // 0x04
	ERR_UNEXPECTED_OP // 0x05
)
const (
	_                = iota
	TYPE_GARAGE_DOOR // 0x01
	TYPE_LIGHT       // 0x02
	TYPE_TEMP        // 0x03
	TYPE_PRESSURE    // 0x04
)
const (
	_          = iota
	STATUS_ON  // 0x01
	STATUS_OFF // 0x02
)

// Full packet
type PDU struct {
	Version    int8   // 1 byte
	Operation  int8   // 1 byte
	Error      int8   // 1 byte
	Reserved   int8   // 1 byte
	Identifier string // 4 bytes
	Devices    [5]Device
}

type Device struct {
	Id     int8    // 1 byte
	Type   int8    // 1 byte
	Status int8    // 1 byte
	Action int8    // 1 bytes
	Value  float32 // 4 bytes
}

func (pdu PDU) Write(conn *tls.Conn) (err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, pdu.Version)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	err = binary.Write(buf, binary.BigEndian, pdu.Operation)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	err = binary.Write(buf, binary.BigEndian, pdu.Error)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	err = binary.Write(buf, binary.BigEndian, pdu.Reserved)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	err = binary.Write(buf, binary.BigEndian, []byte(pdu.Identifier)[0:8])
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	for i := 0; i < 5; i++ {
		err = binary.Write(buf, binary.BigEndian, pdu.Devices[i].Id)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
			return
		}
		err = binary.Write(buf, binary.BigEndian, pdu.Devices[i].Type)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
			return
		}
		err = binary.Write(buf, binary.BigEndian, pdu.Devices[i].Status)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
			return
		}
		err = binary.Write(buf, binary.BigEndian, pdu.Devices[i].Action)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
			return
		}
		err = binary.Write(buf, binary.BigEndian, pdu.Devices[i].Value)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
			return
		}
	}

	fmt.Printf("Marshalled message to send:\n% x\n", buf.Bytes())
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		fmt.Println("ioutil.ReadAll failed:", err)
		return
	}
	fmt.Printf("Sending %d bytes...\n", len(bytes))
	sent, err := conn.Write(bytes)
	fmt.Printf("Sent %d bytes.\n", sent)
	return
}
//read PDU information
func readPDU(conn *tls.Conn) (pdu PDU, err error) {
	var buffer [52]byte
	receivedBytes := 0
	for receivedBytes < 52 {
		temprLen, err := conn.Read(buffer[receivedBytes:])
		receivedBytes += temprLen
		if err == io.EOF {
			return pdu, fmt.Errorf("Received EOF before reading whole packet.")
		}
		if err != nil {
			return pdu, fmt.Errorf("Error reading: %s", err.Error())
		}
	}

	fmt.Printf("Received PDU: %+v\n", buffer)

	err = binary.Read(bytes.NewReader(buffer[0:1]), binary.BigEndian, &pdu.Version)
	if err != nil {
		fmt.Println("binary.Read failed on Version:", err)
		return
	}

	err = binary.Read(bytes.NewReader(buffer[1:2]), binary.BigEndian, &pdu.Operation)
	if err != nil {
		fmt.Println("binary.Read failed on Operation:", err)
		return
	}

	err = binary.Read(bytes.NewReader(buffer[2:3]), binary.BigEndian, &pdu.Error)
	if err != nil {
		fmt.Println("binary.Read failed on Error:", err)
		return
	}

	pdu.Identifier = strings.TrimRight(string(buffer[4:12]), "\x00")

	for i := 1; i <= 5; i++ {
		dev := Device{}
		err = binary.Read(bytes.NewReader(buffer[4+i*8:4+i*8+1]), binary.BigEndian, &dev.Id)
		if err != nil {
			fmt.Println("binary.Read failed on dev[%d].Id Error:", i, err)
			return
		}
		err = binary.Read(bytes.NewReader(buffer[4+i*8+1:4+i*8+2]), binary.BigEndian, &dev.Type)
		if err != nil {
			fmt.Println("binary.Read failed on dev[%d].Id Type:", i, err)
			return
		}
		err = binary.Read(bytes.NewReader(buffer[4+i*8+2:4+i*8+3]), binary.BigEndian, &dev.Status)
		if err != nil {
			fmt.Println("binary.Read failed on dev[%d].Id Status:", i, err)
			return
		}
		err = binary.Read(bytes.NewReader(buffer[4+i*8+3:4+i*8+4]), binary.BigEndian, &dev.Action)
		if err != nil {
			fmt.Println("binary.Read failed on dev[%d].Id Action:", i, err)
			return
		}
		err = binary.Read(bytes.NewReader(buffer[4+i*8+4:4+i*8+8]), binary.BigEndian, &dev.Value)
		if err != nil {
			fmt.Println("binary.Read failed on dev[%d].Id Value:", i, err)
			return
		}
		pdu.Devices[i-1] = dev
	}

	return
}
