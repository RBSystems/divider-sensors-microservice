package main

import (
	"sync"

	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

func main() {
	port := ":7006"
	// router := echo.New()
	//
	// //Functionality endpoints
	//
	// //Status endpoints
	//
	// server := http.Server{
	// 	Addr:           port,
	// 	MaxHeaderBytes: 1024 * 10,
	// }
	//
	// router.StartServer(&server)

	filters := []string{}
	en := eventinfrastructure.NewEventNode("RoomDivide", port, filters, "10.5.34.65:7000")

	var wg sync.WaitGroup
	handlers.StartReading(en, &wg)
	wg.Wait()
}
