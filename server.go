package main

import (
	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"log"
	"os"
	"strings"
)

func main() {
	//Get hostname, building and room.
	//hostname := os.Getenv("PI_HOSTNAME")
	hostname := "ITB-1101-CP5"
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	fmt.Printf("Hostname: %s, Building: %s, Room: %s, Device: %s", hostname, building, room, device)

	filters := []string{eventinfrastructure.RoomDivide}
	en := eventinfrastructure.NewEventNode("Divider Sensors", "7006", filters, os.Getenv("EVENT_ROUTER_ADDRESS"))

	handlers.ReadSensors(hostname, building, room, device, en)
	return
}
