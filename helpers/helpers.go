package helpers

import (
	"log"
	"sync"

	"github.com/byuoitav/common/events"
)

// CONNECTED represents the signal sent by the sensors when the rooms are connected.
const CONNECTED = 1

// DISCONNECTED represents the signal sent by the sensors when the rooms are disconnected.
const DISCONNECTED = 0

// EN is the EventNode object used to publish events.
var EN *events.EventNode

// SetEventNode sets the EventNode object used by the microservice.
func SetEventNode(en *events.EventNode) {
	EN = en
}

// StartReading sets up which pins to read from, and begins reading.
func StartReading(wg *sync.WaitGroup) {
	dc, err := ReadConfig()
	pinList := dc.Pins
	if err != nil {
		log.Printf("Ah dang, I couldn't get the pins...")
		return
	}

	wg.Add(len(pinList))
	for i := range pinList {
		go readSensors(pinList[i], wg)
	}
}

// Connect processes all changes that need to happen when the rooms are connected.
func Connect(p Pin) {
	ConnectedEvent(p)
	DSPChange(p, CONNECTED)
}

// Disconnect processes all changes that need to happen when the rooms are disconnected.
func Disconnect(p Pin) {
	DisconnectedEvent(p)
	DSPChange(p, DISCONNECTED)
}
