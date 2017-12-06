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
	Opened []string `json:"opened,omitempty"`
	Closed []string `json:"closed,omitempty"`
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
			status.Closed = append(status.Closed, msg)
		}
		if state == OPEN {
			msg := fmt.Sprintf("%s", pinList[j].Preset)
			status.Opened = append(status.Opened, msg)
		}
		if state == -1 {
			log.Printf("Cannot read status for pin %s.", pinList[j].Num)
		}
	}
	log.Printf("Success")
	return context.JSON(http.StatusOK, status)
}

// func CheckPinStatus(context echo.Context) error {
// 	dc, err := ReadConfig()
// 	pinList := dc.Pins
// 	number := context.Param("number")
// 	if err != nil {
// 		log.Printf("Couldn't read pins")
// 		return context.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	msg := ""
// 	n := -1
// 	for j := range pinList {
// 		if pinList[j].Num == number {
// 			n = j
// 			state := ReadPinStatus(pinList[j])
// 			if state == CLOSED {
// 				msg = fmt.Sprintf("Rooms %s are disconnected.\n", pinList[j].Preset)
// 			}
// 			if state == OPEN {
// 				msg = fmt.Sprintf("Rooms %s are connected.\n", pinList[j].Preset)
// 			}
// 			if state == -1 {
// 				msg = fmt.Sprintf("Cannot read status for pin %s.\n", pinList[j].Num)
// 				return context.JSON(http.StatusInternalServerError, err.Error())
// 			}
// 		}
// 	}
// 	return context.JSON(http.StatusOK, statusevaluators.Input{: fmt.Sprintf("%s:%s", pinList[n].Num, msg)})
// }

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
