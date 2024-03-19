package main

import (
	"fmt"
	"net/http"
	"todo/internal/controllers"
	"todo/internal/logging"
	"todo/internal/services"

	"github.com/mariusfa/gofl/v2/config"
)

func setup() *http.ServeMux {
	router := http.NewServeMux()
	services := services.New()
	controllers := controllers.New(services)
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
