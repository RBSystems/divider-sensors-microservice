package helpers

import (
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

const CONNECTED = 1
const DISCONNECTED = 0

type DividerConfig struct {
	Pins []Pin `json:"pins"`
}

type Pin struct {
	Num         string `json:"num"`
	Preset      string `json:"preset"`
	DSP         string `json:"dsp"`
	ControlName string `json:"control name"`
}

func Connect(p Pin, en *eventinfrastructure.EventNode) {
	ConnectedEvent(p, en)
	DSPChange(p, CONNECTED)
}

func Disconnect(p Pin, en *eventinfrastructure.EventNode) {
	DisconnectedEvent(p, en)
	DSPChange(p, DISCONNECTED)
}
