package helpers

// DividerConfig contains the information for how each pin is configured.
type DividerConfig struct {
	Pins       []Pin     `json:"pins"`
	Connect    []Request `json:"connect"`
	Disconnect []Request `json:"disconnect"`
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
	Broken       []string    `json:"broken,omitempty"`
	Name         string      `json:"connection name,omitempty"`
	Values       interface{} `json:"values,omitempty"`
}

// Request contains the information as to what information will be pulled in a room
type Request struct {
	Method   string                 `json:"method"`
	Port     string                 `json:"port"`
	Host     string                 `json:"host"`
	Endpoint string                 `json:"endpoint"`
	Body     map[string]interface{} `json:"body"`
}
