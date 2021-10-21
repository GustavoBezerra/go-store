package routes

import (
	"loja/controllers"
	"net/http"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}
