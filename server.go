package main

import (
	"sync"

	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
)

func main() {
	filters := []string{eventinfrastructure.RoomDivide}
	en := eventinfrastructure.NewEventNode("Divider Sensors", "7006", filters, "10.5.34.65:7000")

	var wg sync.WaitGroup
	handlers.DividerStatus(en, &wg)
	wg.Wait()
}
