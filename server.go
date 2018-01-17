package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/divider-sensors-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
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

	//Functionality endpoints

	//Status endpoints
	secure.GET("/status", handlers.AllPinStatus)
	//secure.GET("/nodestatus")

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	filters := []string{}
	en := eventinfrastructure.NewEventNode("RoomDivide", filters, os.Getenv("EVENT_ROUTER_ADDRESS"))

	var wg sync.WaitGroup
	handlers.StartReading(en, &wg)

	router.StartServer(&server)
}
