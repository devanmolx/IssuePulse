package routes

import (
	"backend/controller"
	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/login", controller.Login)
}
