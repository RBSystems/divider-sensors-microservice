package helpers

import (
	"os"
	"strings"
	"time"

	"github.com/byuoitav/common/events"
	"github.com/byuoitav/common/log"
)

// ConnectedEvent builds and publishes an event to the EventRouter.
func ConnectedEvent(p Pin) {
	var CE events.Event
	var con events.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	con.Type = events.DIVISION
	con.Requestor = hostname
	con.EventCause = events.ROOMDIVISION
	con.Device = device
	con.EventInfoKey = "connect"
	con.EventInfoValue = p.Preset

	CE.Hostname = hostname
	CE.Timestamp = time.Now().Format(time.RFC3339)
	CE.LocalEnvironment = true
	CE.Event = con
	CE.Building = building
	CE.Room = room

	log.L.Debugf("Connecting these rooms: %s", p.Preset)
	EN.PublishEvent(events.RoomDivide, CE)
}

// DisconnectedEvent builds and publishes an event to the EventRouter.
func DisconnectedEvent(p Pin) {
	var DE events.Event
	var disc events.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	disc.Type = events.DIVISION
	disc.Requestor = hostname
	disc.EventCause = events.ROOMDIVISION
	disc.Device = device
	disc.EventInfoKey = "disconnect"
	disc.EventInfoValue = p.Preset

	DE.Hostname = hostname
	DE.Timestamp = time.Now().Format(time.RFC3339)
	DE.LocalEnvironment = true
	DE.Event = disc
	DE.Building = building
	DE.Room = room

	log.L.Debugf("Disconnecting these rooms: %s", p.Preset)
	EN.PublishEvent(events.RoomDivide, DE)
}
