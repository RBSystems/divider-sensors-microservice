package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/byuoitav/common/log"

	"net/http"
	"sync"

	"github.com/byuoitav/common/events"
)

// CONNECTED represents the signal sent by the sensors when the rooms are connected.
const CONNECTED = 1

// DISCONNECTED represents the signal sent by the sensors when the rooms are disconnected.
const DISCONNECTED = 0

// EN is the EventNode object used to publish events.
var EN *events.EventNode

var DC DividerConfig

// SetEventNode sets the EventNode object used by the microservice.
func SetEventNode(en *events.EventNode) {
	EN = en
}

// StartReading sets up which pins to read from, and begins reading.
func StartReading(wg *sync.WaitGroup) {
	DC, err := ReadConfig()
	pinList := DC.Pins
	if err != nil {
		log.L.Error("Ah dang, I couldn't get the pins...")
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
	for _, req := range DC.Connect {
		MakeRequest(req)
	}
}

// Disconnect processes all changes that need to happen when the rooms are disconnected.
func Disconnect(p Pin) {
	DisconnectedEvent(p)
	DSPChange(p, DISCONNECTED)
	log.L.Infof("Trying to Disconnect")
	for _, req := range DC.Disconnect {
		MakeRequest(req)
	}
}

// MakeRequest makes a request, WHAAAA????
func MakeRequest(r Request) error {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s:%s/%v/", r.Host, r.Port, r.Endpoint)

	body, err := json.Marshal(r.Body)
	if err != nil {
		log.L.Errorf("Failed to marshal body. ERROR: %s", err.Error())
		return err
	}
	Req, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
	if err != nil {
		log.L.Errorf("Failed to make request. ERROR: %s", err.Error())
		return err
	}

	Resp, err := client.Do(Req)
	if err != nil {
		log.L.Errorf("Failed to send request. ERROR: %s", err.Error())
		return err
	}

	if Resp.StatusCode/100 != 2 {
		log.L.Errorf("NON 200 RESPONSE!!!. ERROR: %s", err.Error())
		return err
	}

	return nil
}
