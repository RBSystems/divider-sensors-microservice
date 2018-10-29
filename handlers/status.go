package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/byuoitav/common/log"
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
		log.L.Errorf("Couldn't read pins")
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

	status.Values = helpers.Messenger.GetState()

	log.L.Debugf("Success")
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

// PresetForHostname returns the current preset a specific hostname should be on
func PresetForHostname(context echo.Context) error {
	hostname := context.Param("hostname")

	if len(hostname) == 0 {
		return errors.New("must include a hostname")
	}

	dc, err := helpers.ReadConfig()
	pinList := dc.Pins
	if err != nil {
		log.L.Errorf("Couldn't read pins")
		return context.JSON(http.StatusInternalServerError, err)
	}

	// this endpoint is currently only supported on rooms with one pin
	if len(pinList) == 0 || len(pinList) > 1 {
		return context.JSON(http.StatusInternalServerError, fmt.Sprintf("not suppored in this room"))
	}

	state := ReadPinStatus(pinList[0])
	if state == helpers.CONNECTED {
		for _, connectEvent := range dc.ConnectEvents {
			if strings.EqualFold(connectEvent.TargetDevice.DeviceID, hostname) {
				return context.JSON(http.StatusOK, connectEvent.Value)
			}
		}
		return context.JSON(http.StatusInternalServerError, "Pins are connected, but no preset was found for this hostname.")
	} else if state == helpers.DISCONNECTED {
		for _, disconnectEvent := range dc.DisconnectEvents {
			if strings.EqualFold(disconnectEvent.TargetDevice.DeviceID, hostname) {
				return context.JSON(http.StatusOK, disconnectEvent.Value)
			}
		}
		return context.JSON(http.StatusInternalServerError, "Pins are disconnected, but no preset was found for this hostname.")
	}

	return context.JSON(http.StatusInternalServerError, fmt.Sprintf("cannot read status for pin %v.", pinList[0].Num))
}
