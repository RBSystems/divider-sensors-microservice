package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// AllPinStatus builds the Status object with the information about the state of the dividers.
func AllPinStatus(context echo.Context) error {
	dc, err := helpers.ReadConfig()
	pinList := dc.Pins
	status := helpers.Status{}
	if err != nil {
		log.Printf("Couldn't read pins")
		return context.JSON(http.StatusInternalServerError, err.Error())
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
			msg := fmt.Sprintf("Cannot read status for pin %s.", pinList[j].Num)
			status.Broken = append(status.Broken, msg)
		}
	}

	status.Name, status.Values = helpers.EN.Node.GetState()

	log.Printf("Success")
	return context.JSON(http.StatusOK, status)
}

// ReadPinStatus reads the status for an individual pin.
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
