/**
CS 544 Computer Networks
6-1-2016

Group 2:
	Daniel Speichert
	Kenneth Balogh
	Arudra Venkat
	Xiaofeng Zhou

Purpose:
	SERVICE
	data.go defines fake 'static' information that is served over MCGP protocol.
    This would normally be dynamically harvested from connected devices via
    other means.
*/

package main

import "math/rand"

//default devices used for server implementation and demo
var devices []Device = []Device{
	Device{Id: 1, Type: TYPE_GARAGE_DOOR, Status: STATUS_ON},
	Device{Id: 2, Type: TYPE_LIGHT, Status: STATUS_ON},
	Device{Id: 3, Type: TYPE_LIGHT, Status: STATUS_ON},
	Device{Id: 4, Type: TYPE_LIGHT, Status: STATUS_OFF},
	Device{Id: 5, Type: TYPE_TEMP, Status: STATUS_ON, Value: rand.Float32() * 10},
	Device{Id: 6, Type: TYPE_TEMP, Status: STATUS_OFF},
	Device{Id: 7, Type: TYPE_PRESSURE, Status: STATUS_ON, Value: rand.Float32() * 10},
}
