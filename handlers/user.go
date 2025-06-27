package handlers

import (
	"encoding/json"
	"net/http"

	"custom-api-server/db"
	"custom-api-server/models"
	"custom-api-server/utils"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()

	result := db.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.ID == uuid.Nil {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(input.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id = ?", input.ID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	db.DB.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID uuid.UUID `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.ID == uuid.Nil {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	result := db.DB.Delete(&models.User{}, "id = ?", input.ID)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "User not found or could not be deleted", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
