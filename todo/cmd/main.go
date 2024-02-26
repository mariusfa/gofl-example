package main

import (
	"fmt"
	"net/http"
	"todo/app"
	"todo/internal/domain"

	"github.com/mariusfa/gofl/config"
)

func setup() *http.ServeMux {
	router := http.NewServeMux()
	services := domain.NewServices()
	controllers := app.NewControllers(services)
	controllers.RegisterRoutes(router)

	return router
}

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%d", config.Port)
	router := setup()

	println("Listening on", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
