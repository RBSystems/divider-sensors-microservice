package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/common/events"
	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/divider-sensors-microservice/helpers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
	helpers.SetEventNode(events.NewEventNode("RoomDivide", os.Getenv("EVENT_ROUTER_ADDRESS"), filters))

	//Status endpoints
	secure.GET("/status", handlers.AllPinStatus)

	var wg sync.WaitGroup
	helpers.StartReading(&wg)

	router.StartServer(&server)
}
