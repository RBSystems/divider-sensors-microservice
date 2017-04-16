package main

import (
	"time"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"fmt"
)

func main() {
	r := raspi.NewAdaptor()
	sensor := gpio.NewDirectPinDriver(r, "7")

	TIMES := 0

	for TIMES < 200 {
		time.sleep(0.5)
		read, err := sensor.DigitalRead()
		if read == 1 {
			Printf("read 1\n")
		}
		if read == 0 {
			Printf("read 0\n")
		}
		if err != nil {
			Printf("Something went wrong\n")
		}
		TIMES += 1
	}
}
