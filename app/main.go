package main

import (
	"net/http"
	"web-service/app/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
