package main

import (
	"net/http"
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
	secure.GET("/allstatus", handlers.AllPinStatus)
	//secure.GET("/status/:number", handlers.CheckPinStatus)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	filters := []string{}
	en := eventinfrastructure.NewEventNode("RoomDivide", "7006", filters, "10.5.34.65:7000")

	var wg sync.WaitGroup
	handlers.StartReading(en, &wg)

	router.StartServer(&server)
}
