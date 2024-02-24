package main

import (
	"net/http"
	"todo/app"
)

func main() {
	router := http.NewServeMux()
	controllers := app.NewControllers()
	controllers.RegisterRoutes(router)

	http.ListenAndServe(":8080", router)
}