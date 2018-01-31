package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const COUNTER_MAX = 10

func StartReading(en *eventinfrastructure.EventNode, wg *sync.WaitGroup) {
	dc, err := ReadConfig()
	pinList := dc.Pins
	if err != nil {
		log.Printf("Ah dang, I couldn't get the pins...")
		return
	}

	wg.Add(len(pinList))
	for i := range pinList {
		go ReadSensors(pinList[i], en, wg)
	}
}

func ReadConfig() (helpers.DividerConfig, error) {
	CONFIG := os.Getenv("CONTACTS_CONFIG_FILE")
	body, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		log.Printf("Failed to read body from file %s: %s", CONFIG, err)
		return helpers.DividerConfig{}, err
	}

	var config helpers.DividerConfig
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Printf("Failed to unmarshal body from file %s: %s", CONFIG, err)
		return helpers.DividerConfig{}, err
	}

	return config, nil
}

func ReadSensors(p helpers.Pin, en *eventinfrastructure.EventNode, wg *sync.WaitGroup) {
	defer wg.Done()
	//Establish connection to the GPIO
	r := raspi.NewAdaptor()
	sensor := gpio.NewDirectPinDriver(r, p.Num)
	read, err := sensor.DigitalRead()

	//Initialize counter variables
	times := 0
	connectedCount := 0
	disconnectedCount := 0
	curState := read
	for {
		//Based on the time.Sleep for .1 seconds, after 300 loops it will have been 5 min.
		for times < 300 {
			//Read at every interval to assess a status change
			time.Sleep(100 * time.Millisecond)
			read, err = sensor.DigitalRead()

			if read != curState {
				//Dividers read as open
				if read == helpers.CONNECTED {
					connectedCount += 1
					disconnectedCount = 0
					if connectedCount == COUNTER_MAX {
						//Send open event
						curState = helpers.CONNECTED
						helpers.Connect(p, en)
					}
				}

				//Dividers read as closed
				if read == helpers.DISCONNECTED {
					disconnectedCount += 1
					connectedCount = 0
					if disconnectedCount == COUNTER_MAX {
						//Send closed event
						curState = helpers.DISCONNECTED
						helpers.Disconnect(p, en)
					}
				}
				if err != nil {
					log.Printf("Something went wrong\n")
				}
			}
			times += 1
		}
		//Every 5 min, send out an event of the current state.
		read, err = sensor.DigitalRead()
		if read == helpers.CONNECTED {
			curState = helpers.CONNECTED
			helpers.Connect(p, en)
		}
		if read == helpers.DISCONNECTED {
			curState = helpers.DISCONNECTED
			helpers.Disconnect(p, en)
		}
		connectedCount = 0
		disconnectedCount = 0
		times = 0
	}
}
