package handlers

import (
	"fmt"
	"log"

	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type Status struct {
	Disconnected []string    `json:"disconnected,omitempty"`
	Connected    []string    `json:"connected,omitempty"`
	Name         string      `json:"connection name,omitempty"`
	Values       interface{} `json:"values,omitempty"`
}

func AllPinStatus(en *eventinfrastructure.EventNode) Status {
	dc, err := ReadConfig()
	pinList := dc.Pins
	status := Status{}
	if err != nil {
		log.Printf("Couldn't read pins")
		return status
	}

	for j := range pinList {
		state := ReadPinStatus(pinList[j])
		if state == helpers.CONNECTED {
			msg := fmt.Sprintf("%s", pinList[j].Preset)
			status.Connected = append(status.Connected, msg)
		}
		if state == helpers.DISCONNECTED {
			msg := fmt.Sprintf("%s", pinList[j].Preset)
			status.Disconnected = append(status.Disconnected, msg)
		}
		if state == -1 {
			log.Printf("Cannot read status for pin %s.", pinList[j].Num)
		}
	}

	status.Name, status.Values = en.Node.GetState()

	log.Printf("Success")
	return status
}

func ReadPinStatus(p helpers.Pin) int {
	//Establish connection to the GPIO
	r := raspi.NewAdaptor()
	sensor := gpio.NewDirectPinDriver(r, p.Num)
	read, err := sensor.DigitalRead()
	if err != nil {
		return -1
	}
	return read
}
