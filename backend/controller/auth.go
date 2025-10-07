package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var users []models.User
var userID = 1

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	req.ID = userID
	userID++
	req.Password = hashedPassword
	users = append(users, req)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	for _, user := range users {
		if user.Username == req.Username {
			if utils.CheckPasswordHash(req.Password, user.Password) {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
				return
			}
			break
		}
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)

	fmt.Println("Current users:", users)
}
