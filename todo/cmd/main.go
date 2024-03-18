package main

import (
	"fmt"
	"net/http"
	"todo/app"
	"todo/internal"
	"todo/internal/logging"

	"github.com/mariusfa/gofl/v2/config"
)

func setup() *http.ServeMux {
	router := http.NewServeMux()
	services := internal.NewServices()
	controllers := app.NewControllers(services)
	controllers.RegisterRoutes(router)

	return router
}

type Config struct {
	Port string
}

func main() {
	logging.SetupAppLogger("todo")

	var appConfig Config
	err := config.GetConfig(".env", &appConfig)
	if err != nil {
		panic(err)
	}

	router := setup()

	addr := fmt.Sprintf(":%s", appConfig.Port)
	logging.AppLogger.Info("Starting server on " + addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
