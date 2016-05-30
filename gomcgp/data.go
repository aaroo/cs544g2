package main

import "math/rand"

var devices []Device = []Device{
	Device{Id: 1, Type: TYPE_GARAGE_DOOR, Status: STATUS_ON},
	Device{Id: 2, Type: TYPE_LIGHT, Status: STATUS_ON},
	Device{Id: 3, Type: TYPE_LIGHT, Status: STATUS_ON},
	Device{Id: 4, Type: TYPE_LIGHT, Status: STATUS_OFF},
	Device{Id: 5, Type: TYPE_TEMP, Status: STATUS_ON, Value: rand.Float32() * 10},
	Device{Id: 6, Type: TYPE_TEMP, Status: STATUS_OFF},
	Device{Id: 7, Type: TYPE_PRESSURE, Status: STATUS_ON, Value: rand.Float32() * 10},
}
