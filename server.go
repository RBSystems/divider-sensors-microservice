package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var en *eventinfrastructure.EventNode

func main() {
	port := ":8200"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	filters := []string{}
	helpers.SetEventNode(eventinfrastructure.NewEventNode("RoomDivide", filters, os.Getenv("EVENT_ROUTER_ADDRESS")))

	//Status endpoints
	secure.GET("/status", state)

	var wg sync.WaitGroup
	helpers.StartReading(&wg)

	router.StartServer(&server)
}

func state(context echo.Context) error {
	var status helpers.Status = handlers.AllPinStatus(en)
	return context.JSON(http.StatusOK, status)
}
