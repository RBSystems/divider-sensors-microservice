package helpers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func DSPChange(p Pin, state int) {
	client := &http.Client{}
	//I put in a request bin url for testing since the DSP is in another room, but this will change back to having a DSP address.
	url := fmt.Sprintf("http://linux-knight.byu.edu:8016/%v/generic/%v/%v", p.DSP, p.ControlName, state)
	//url := fmt.Sprintf("https://requestb.in/1hyeg4j1")

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte{}))
	if err != nil {
		// handle error
		log.Printf("%s", err)
	}
	_, err = client.Do(req)
	return
}
