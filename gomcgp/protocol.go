package main

const (
	SUPPORTED_VERSION = 1

	_                 = iota
	OP_VERSION        // 0x01
	OP_AUTHENTICATE   // 0x02
	OP_LIST           // 0x03
	OP_LIST_CONTINUED // 0x04
	OP_CONTROL        // 0x05

	_                = iota
	ERR_VERSION      // 0x01
	ERR_AUTHENTICATE // 0x02
	ERR_LIST         // 0x03
	ERR_CONTROL      // 0x04
)

// Full packet
type PDU struct {
	Version    int8   // 1 byte
	Operation  int8   // 1 byte
	Error      int8   // 1 byte
	Identifier string // 8 bytes
	Devices    [5]Device
}

type Device struct {
	Id     int8    // 1 byte
	Type   int8    // 1 byte
	Status int8    // 1 byte
	Action int8    // 1 bytes
	Value  float32 // 4 bytes
}
