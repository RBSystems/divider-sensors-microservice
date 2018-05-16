package helpers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

// DSPChange sends a request to the DSP Microservice for the room to make the necessary audio changes.
func DSPChange(p Pin, state int) {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s:8016/%v/generic/%v/%v", os.Getenv("DSP_MICROSERVICE_ADDRESS"), p.DSP, p.ControlName, state)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte{}))
	if err != nil {
		// handle error
		log.Printf("%v", err.Error())
	}
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Printf("%v", err.Error())
	}
	return
}
