package helpers

// DividerConfig contains the information for how each pin is configured.
type DividerConfig struct {
	Pins []Pin `json:"pins"`
}

// Pin lists the configuration for this specific pin.
type Pin struct {
	Num         string `json:"num"`
	Preset      string `json:"preset"`
	DSP         string `json:"dsp"`
	ControlName string `json:"control name"`
}

// Status contains the information to return when asked for the state of the dividers.
type Status struct {
	Disconnected []string    `json:"disconnected,omitempty"`
	Connected    []string    `json:"connected,omitempty"`
	Name         string      `json:"connection name,omitempty"`
	Values       interface{} `json:"values,omitempty"`
}
