package events

import (
	"log"
	"strings"
	"time"

	ei "github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

func OpenedEvent(preset string, en *ei.EventNode) {
	var OE ei.Event
	var open ei.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := "ITB-1101-CP5"
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	open.Type = ei.DIVISION
	open.Requestor = hostname
	open.EventCause = ei.ROOMDIVISION
	open.Device = device
	open.EventInfoKey = "connect"
	open.EventInfoValue = preset

	log.Printf("%s", open.EventInfoValue)

	OE.Hostname = hostname
	OE.Timestamp = time.Now().Format(time.RFC3339)
	OE.LocalEnvironment = true
	OE.Event = open
	OE.Building = building
	OE.Room = room

	log.Printf("OPEN EVENT")
	en.PublishEvent(OE, "DIVISION")
}

func ClosedEvent(preset string, en *ei.EventNode) {
	var CE ei.Event
	var closed ei.EventInfo

	//Get Hostname, Building, Room and Device
	hostname := "ITB-1101-CP5"
	roomInfo := strings.Split(hostname, "-")
	building := roomInfo[0]
	room := roomInfo[1]
	device := roomInfo[2]

	closed.Type = ei.DIVISION
	closed.Requestor = hostname
	closed.EventCause = ei.ROOMDIVISION
	closed.Device = device
	closed.EventInfoKey = "disconnect"
	closed.EventInfoValue = preset

	log.Printf("%s", closed.EventInfoValue)

	CE.Hostname = hostname
	CE.Timestamp = time.Now().Format(time.RFC3339)
	CE.LocalEnvironment = true
	CE.Event = closed
	CE.Building = building
	CE.Room = room

	log.Printf("CLOSED EVENT")
	en.PublishEvent(CE, "DIVISION")
}
