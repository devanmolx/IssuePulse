package main

import (
	"backend/db"
	"backend/models"
	"backend/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	if err := db.Connect(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	db.DB.AutoMigrate(&models.User{})

	routes.AuthRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting at http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
