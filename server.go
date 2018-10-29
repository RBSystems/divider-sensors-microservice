package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/central-event-system/hub/base"
	"github.com/byuoitav/central-event-system/messenger"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
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

	var err *nerr.E
	helpers.Messenger, err = messenger.BuildMessenger(os.Getenv("HUB_ADDRESS"), base.Messenger, 1000)
	if err != nil {
		log.L.Errorf("there was an error building the messenger: ", err.String())
		return
	}

	//Status endpoints
	secure.GET("/status", handlers.AllPinStatus)
	secure.GET("/preset/:hostname", handlers.PresetForHostname)

	var wg sync.WaitGroup
	helpers.StartReading(&wg)

	router.StartServer(&server)
}
