package helpers

import (
	"log"
	"os"
	"strings"
	"time"

	ei "github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

// ConnectedEvent builds and publishes an event to the EventRouter.
func ConnectedEvent(p Pin) {
	var CE ei.Event
	var con ei.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	con.Type = ei.DIVISION
	con.Requestor = hostname
	con.EventCause = ei.ROOMDIVISION
	con.Device = device
	con.EventInfoKey = "connect"
	con.EventInfoValue = p.Preset

	log.Printf("%s", con.EventInfoValue)

	CE.Hostname = hostname
	CE.Timestamp = time.Now().Format(time.RFC3339)
	CE.LocalEnvironment = true
	CE.Event = con
	CE.Building = building
	CE.Room = room

	log.Printf("Connecting these rooms: %s", p.Preset)
	EN.PublishEvent(CE, ei.RoomDivide)
}

// DisconnectedEvent builds and publishes an event to the EventRouter.
func DisconnectedEvent(p Pin) {
	var DE ei.Event
	var disc ei.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	disc.Type = ei.DIVISION
	disc.Requestor = hostname
	disc.EventCause = ei.ROOMDIVISION
	disc.Device = device
	disc.EventInfoKey = "disconnect"
	disc.EventInfoValue = p.Preset

	log.Printf("%s", disc.EventInfoValue)

	DE.Hostname = hostname
	DE.Timestamp = time.Now().Format(time.RFC3339)
	DE.LocalEnvironment = true
	DE.Event = disc
	DE.Building = building
	DE.Room = room

	log.Printf("Disconnecting these rooms: %s", p.Preset)
	EN.PublishEvent(DE, ei.RoomDivide)
}
