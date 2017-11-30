package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/byuoitav/divider-sensors-microservice/events"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const COUNTER_MAX = 10

const CONFIG = "config.json"

type DividerConfig struct {
	Pins []Pin `json:"pins"`
}

type Pin struct {
	Num    string `json:"num"`
	Preset string `json:"preset"`
}

func DividerStatus(en *eventinfrastructure.EventNode, wg *sync.WaitGroup) {
	dc, err := ReadConfig()
	pinList := dc.Pins
	if err != nil {
		log.Printf("Aww dang, I couldn't get the pins...")
		return
	}

	wg.Add(len(pinList))
	for i := range pinList {
		go ReadSensors(pinList[i].Num, pinList[i].Preset, en, wg)
	}
}

func ReadConfig() (DividerConfig, error) {
	body, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		log.Printf("Failed to read body from file %s: %s", CONFIG, err)
		return DividerConfig{}, err
	}

	var config DividerConfig
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Printf("Failed to unmarshal body from file %s: %s", CONFIG, err)
		return DividerConfig{}, err
	}

	return config, nil
}

func ReadSensors(pin_num string, preset string, en *eventinfrastructure.EventNode, wg *sync.WaitGroup) {
	defer wg.Done()
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
					if open_count == COUNTER_MAX {
						//Send open event
						cur_state = 1
						events.OpenedEvent(preset, en)
					}
				}

				//Dividers read as closed
				if read == 0 {
					closed_count += 1
					open_count = 0
					if closed_count == COUNTER_MAX {
						//Send closed event
						cur_state = 0
						events.ClosedEvent(preset, en)
					}
				}
				if err != nil {
					log.Printf("Something went wrong\n")
				}
			}
			times += 1
		}
		open_count = 0
		closed_count = 0
		times = 0
	}
}
