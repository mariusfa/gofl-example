package main

import (
	"net/http"
	"todo/app"
)

func setup() *http.ServeMux {
	router := http.NewServeMux()
	controllers := app.NewControllers()
	controllers.RegisterRoutes(router)

	return router
}

func main() {
	router := setup()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}