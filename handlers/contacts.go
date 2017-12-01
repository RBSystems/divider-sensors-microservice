package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const COUNTER_MAX = 10
const OPEN = helpers.OPEN
const CLOSED = helpers.CLOSED

const CONFIG = "config.json"

func StartReading(en *eventinfrastructure.EventNode, wg *sync.WaitGroup) {
	dc, err := ReadConfig()
	pinList := dc.Pins
	if err != nil {
		log.Printf("Aww dang, I couldn't get the pins...")
		return
	}

	wg.Add(len(pinList))
	for i := range pinList {
		go ReadSensors(pinList[i], en, wg)
	}
}

func ReadConfig() (helpers.DividerConfig, error) {
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
	openCount := 0
	closedCount := 0
	curState := read
	for {
		for times < 50 {
			//Read at every interval to assess a status change
			time.Sleep(100 * time.Millisecond)
			read, err = sensor.DigitalRead()

			if read != curState {
				//Dividers read as open
				if read == OPEN {
					openCount += 1
					closedCount = 0
					if openCount == COUNTER_MAX {
						//Send open event
						curState = OPEN
						helpers.Disconnect(p, en)
					}
				}

				//Dividers read as closed
				if read == CLOSED {
					closedCount += 1
					openCount = 0
					if closedCount == COUNTER_MAX {
						//Send closed event
						curState = CLOSED
						helpers.Connect(p, en)
					}
				}
				if err != nil {
					log.Printf("Something went wrong\n")
				}
			}
			times += 1
		}
		openCount = 0
		closedCount = 0
		times = 0
	}
}

func AllPinStatus() {
	// dc, err := ReadConfig()
	// pinList := dc.Pins
	// if err != nil {
	// 	log.Printf("Couldn't read pins")
	// 	return
	// }

	// for j := range pinList {
	// 	state := CheckPinStatus(pinList[j])
	// 	if state == CLOSED {
	// 		//Something will happen here. Preset and disconnected?
	// 	}
	// 	else if state == OPEN {
	// 		//Preset and connected?
	// 	}
	// }
}

// func CheckPinStatus(p helper.Pin) {
//
// }
