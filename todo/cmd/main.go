package main

import (
	"fmt"
	"net/http"
	"todo/app"
	"todo/internal/domain"
	"todo/internal/logging"

	"github.com/mariusfa/gofl/v2/config"
)

func setup() *http.ServeMux {
	router := http.NewServeMux()
	services := domain.NewServices()
	controllers := app.NewControllers(services)
	controllers.RegisterRoutes(router)

	return router
}

func main() {
	logging.SetupAppLogger("todo")

	config, err := config.GetConfig(logging.AppLogger)
	if err != nil {
		panic(err)
	}

	router := setup()

	addr := fmt.Sprintf(":%d", config.Port)
	logging.AppLogger.Info("Starting server on " + addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
