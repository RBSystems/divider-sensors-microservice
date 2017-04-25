package events

import (
	"log"
	"time"
	ei "github.com/byuoitav/event-router-microservice/eventinfrastructure"
)
func OpenedEvent(hostname string, building string, room string, device string, en *ei.EventNode) {
	var OE ei.Event
	var open ei.EventInfo

	open.Type = ei.DIVISION
	open.Requestor = hostname
	open.EventCause = ei.ROOMDIVISION
	open.Device = device
	open.EventInfoKey = "open"
	open.EventInfoValue = "open"

	OE.Hostname = hostname
	OE.Timestamp = time.Now().Format(time.RFC3339)
	OE.LocalEnvironment = true
	OE.Event = open
	OE.Building = building
	OE.Room = room

	log.Printf("OPEN EVENT")
}

func ClosedEvent(hostname string, building string, room string, device string, en *ei.EventNode) {
	var CE ei.Event
	var closed ei.EventInfo

	closed.Type = ei.DIVISION
	closed.Requestor = hostname
	closed.EventCause = ei.ROOMDIVISION
	closed.Device = device
	closed.EventInfoKey = "close"
	closed.EventInfoValue = "closed"

	CE.Hostname = hostname
	CE.Timestamp = time.Now().Format(time.RFC3339)
	CE.LocalEnvironment = true
	CE.Event = closed
	CE.Building = building
	CE.Room = room

	log.Printf("CLOSED EVENT")

}


