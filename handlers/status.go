package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/labstack/echo"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type Status struct {
	Disconnected []string `json:"disconnected,omitempty"`
	Connected    []string `json:"connected,omitempty"`
}

func AllPinStatus(context echo.Context) error {
	dc, err := ReadConfig()
	pinList := dc.Pins
	status := Status{}
	if err != nil {
		log.Printf("Couldn't read pins")
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	for j := range pinList {
		state := ReadPinStatus(pinList[j])
		if state == CLOSED {
			msg := fmt.Sprintf("%s", pinList[j].Preset)
			status.Connected = append(status.Connected, msg)
		}
		if state == OPEN {
			msg := fmt.Sprintf("%s", pinList[j].Preset)
			status.Disconnected = append(status.Disconnected, msg)
		}
		if state == -1 {
			log.Printf("Cannot read status for pin %s.", pinList[j].Num)
		}
	}
	log.Printf("Success")
	return context.JSON(http.StatusOK, status)
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
