package main

import (
	"loja/routes"
	"net/http"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
