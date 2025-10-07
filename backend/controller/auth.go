package controller

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

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

	req.Password = hashedPassword

	result := db.DB.Create(&req)
	if result.Error != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	var user models.User

	result := db.DB.Where("username = ?", req.Username).First(&user)

	if result.Error != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})

}
