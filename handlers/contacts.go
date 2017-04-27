package handlers

import (
	"github.com/byuoitav/divider-sensors-microservice/events"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"log"
	"time"
)

const COUNTER_MAX = 10

func ReadSensors(hostname string, building string, room string, device string, pin_num string, preset string, en *eventinfrastructure.EventNode) {
	//Establish connection to the GPIO
	r := raspi.NewAdaptor()
	sensor := gpio.NewDirectPinDriver(r, pin_num)
	read, err := sensor.DigitalRead()

	//Initialize counter variables
	times := 0
	open_count := 0
	closed_count := 0
	cur_state := read

	for {
		for times < 50 {
			//Read at every interval to assess a status change
			time.Sleep(100 * time.Millisecond)
			read, err = sensor.DigitalRead()

			if read != cur_state {
				//Dividers read as open
				if read == 1 {
					open_count += 1
					closed_count = 0
					log.Printf("Open")
					if open_count == COUNTER_MAX {
						//Send open event
						cur_state = 1
						log.Printf("Open Max\n")
						events.OpenedEvent(hostname, building, room, device, preset, en)
					}
				}

				//Dividers read as closed
				if read == 0 {
					closed_count += 1
					open_count = 0
					log.Printf("Close")
					if closed_count == COUNTER_MAX {
						//Send closed event
						cur_state = 0
						log.Printf("Closed max\n")
						events.ClosedEvent(hostname, building, room, device, preset, en)
					}
				}
				if err != nil {
					log.Printf("Something went wrong\n")
				}
			}
			//log.Printf("/")
			times += 1
		}
		open_count = 0
		closed_count = 0
		times = 0
	}
}
