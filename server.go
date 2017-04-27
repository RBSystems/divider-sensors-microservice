package main

import (
	"encoding/json"
	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const CONFIG = "config.json"

type Pins struct {
	Num    string `json:"num"`
	Preset string `json:"preset"`
}

//func toString(p Pin) string {
//	bytes, err := json.Marshal(p)
//	if err != nil {
//		log.Printf(err.Error())
//	}
//	return string(bytes)
//}

func getPins() (Pins, error) {
	body, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		log.Printf("Failed to read body from file %s: %s", CONFIG, err)
		return Pins{}, err
	}

	var config Pins
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Printf("Failed to unmarshal body from file %s: %s", CONFIG, err)
		return Pins{}, err
	}

	return config, nil
}

func main() {
	//Set pin and preset by reading config file
	pin, _ := getPins()
	pin_num := pin.Num
	preset := pin.Preset
	log.Printf("%s and %s", pin_num, preset)
	//Get hostname, building and room.
	//hostname := os.Getenv("PI_HOSTNAME")
	hostname := "ITB-1101-CP5"
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	log.Printf("Hostname: %s, Building: %s, Room: %s, Device: %s", hostname, building, room, device)

	filters := []string{eventinfrastructure.RoomDivide}
	en := eventinfrastructure.NewEventNode("Divider Sensors", "7006", filters, os.Getenv("EVENT_ROUTER_ADDRESS"))

	handlers.ReadSensors(hostname, building, room, device, pin_num, preset, en)
	return
}
