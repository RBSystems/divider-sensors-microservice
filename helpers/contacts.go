package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// CounterMax is the maximum number of times that the sensors should read the same number before sending an event.
const CounterMax = 10

// ReadConfig reads the file that determines the configuration for each pin to be read from.
func ReadConfig() (DividerConfig, error) {
	CONFIG := os.Getenv("CONTACTS_CONFIG_FILE")

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

func readSensors(p Pin, wg *sync.WaitGroup) {
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

	// Endlessly read
	for {
		for times < 50 {
			//Read at every interval to assess a status change
			time.Sleep(100 * time.Millisecond)
			read, err = sensor.DigitalRead()

			if read != curState {
				//Dividers read as open
				if read == CONNECTED {

					connectedCount++
					disconnectedCount = 0

					if connectedCount == CounterMax {
						//Send open event
						curState = CONNECTED
						go Connect(p)
					}
				}

				//Dividers read as closed
				if read == DISCONNECTED {

					disconnectedCount++
					connectedCount = 0

					if disconnectedCount == CounterMax {
						//Send closed event
						curState = DISCONNECTED
						go Disconnect(p)
					}
				}
				if err != nil {
					log.Printf("Something went wrong\n")
				}
			}
			times++
		}
		// Reset counters
		connectedCount = 0
		disconnectedCount = 0
		times = 0
	}
}
