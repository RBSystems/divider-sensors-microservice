package helpers

import (
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

const CLOSED = 0
const OPEN = 1

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
	ClosedEvent(p, en)
	DSPChange(p, CLOSED)
}

func Disconnect(p Pin, en *eventinfrastructure.EventNode) {
	OpenedEvent(p, en)
	DSPChange(p, OPEN)
}
