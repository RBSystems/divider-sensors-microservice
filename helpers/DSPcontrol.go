package helpers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func DSPChange(p Pin, state int) {
	client := &http.Client{}
	url := fmt.Sprintf("http://linux-knight.byu.edu:8016/%v/generic/%v/%v", p.DSP, p.ControlName, state)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte{}))
	if err != nil {
		// handle error
		log.Printf("%s", err)
	}
	_, err = client.Do(req)
	return
}
