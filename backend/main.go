package main

import (
	"backend/routes"
	"fmt"
	"net/http"
)

func main() {

	routes.AuthRoutes()

	fmt.Println("Server starting at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
