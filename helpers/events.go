package helpers

import (
	"os"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/events"
)

// ConnectedEvent builds and publishes an event to the EventRouter.
func ConnectedEvent(p Pin) {
	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]

	roomStuff := events.BasicRoomInfo{
		BuildingID: building,
		RoomID:     room,
	}

	deviceInfo := events.BasicDeviceInfo{
		BasicRoomInfo: roomStuff,
		DeviceID:      hostname,
	}

	event := events.Event{
		GeneratingSystem: hostname,
		Timestamp:        time.Now(),
		AffectedRoom:     roomStuff,
		TargetDevice:     deviceInfo,
		Key:              "connect",
		Value:            p.Preset,
		User:             hostname,
	}

	event.EventTags = append(event.EventTags, events.RoomDivide, "LocalEnv")

	log.L.Debugf("Connecting these rooms: %s", p.Preset)
	EN.PublishEvent(events.RoomDivide, event)
}

// DisconnectedEvent builds and publishes an event to the EventRouter.
func DisconnectedEvent(p Pin) {
	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("PI_HOSTNAME")
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]

	roomStuff := events.BasicRoomInfo{
		BuildingID: building,
		RoomID:     room,
	}

	deviceInfo := events.BasicDeviceInfo{
		BasicRoomInfo: roomStuff,
		DeviceID:      hostname,
	}

	event := events.Event{
		GeneratingSystem: hostname,
		Timestamp:        time.Now(),
		AffectedRoom:     roomStuff,
		TargetDevice:     deviceInfo,
		Key:              "disconnect",
		Value:            p.Preset,
		User:             hostname,
	}

	log.L.Debugf("Disconnecting these rooms: %s", p.Preset)
	EN.PublishEvent(events.RoomDivide, event)
}
