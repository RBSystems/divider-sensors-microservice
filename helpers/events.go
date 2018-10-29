package helpers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/events"
)

// ConnectedEvent builds and publishes an event to the EventRouter.
func ConnectedEvent(p Pin) {
	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("SYSTEM_ID")
	roomInfo := strings.Split(hostname, "-")
	roomID := fmt.Sprintf("%s-%s", roomInfo[0], roomInfo[1])

	roomStuff := events.GenerateBasicRoomInfo(roomID)

	deviceInfo := events.GenerateBasicDeviceInfo(hostname)

	event := events.Event{
		GeneratingSystem: hostname,
		Timestamp:        time.Now(),
		AffectedRoom:     roomStuff,
		TargetDevice:     deviceInfo,
		Key:              "connect",
		Value:            p.Preset,
		User:             hostname,
	}

	event.AddToTags("LocalEnv", events.RoomDivide)

	log.L.Debugf("Connecting these rooms: %s", p.Preset)
	SendEvent(event)
}

// DisconnectedEvent builds and publishes an event to the EventRouter.
func DisconnectedEvent(p Pin) {
	//Get Hostname, Building, Room and Device
	hostname := os.Getenv("SYSTEM_ID")
	roomInfo := strings.Split(hostname, "-")
	roomID := fmt.Sprintf("%s-%s", roomInfo[0], roomInfo[1])

	roomStuff := events.GenerateBasicRoomInfo(roomID)

	deviceInfo := events.GenerateBasicDeviceInfo(hostname)

	event := events.Event{
		GeneratingSystem: hostname,
		Timestamp:        time.Now(),
		AffectedRoom:     roomStuff,
		TargetDevice:     deviceInfo,
		Key:              "disconnect",
		Value:            p.Preset,
		User:             hostname,
	}

	event.AddToTags("LocalEnv", events.RoomDivide)

	log.L.Debugf("Disconnecting these rooms: %s", p.Preset)
	SendEvent(event)
}
