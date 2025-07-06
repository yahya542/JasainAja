package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"jasainaja-backend/database"
	"jasainaja-backend/models"
)

func RegisterProvider(w http.ResponseWriter, r *http.Request) {
	var provider models.Provider
	err := json.NewDecoder(r.Body).Decode(&provider)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO providers (name, email, password) VALUES ($1, $2, $3)`
	_, err = database.DB.Exec(query, provider.Name, provider.Email, provider.Password)
	if err != nil {
		log.Println("❌ Failed to register provider:", err)
		http.Error(w, "Failed to register provider", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Provider registered successfully"})
}

func LoginProvider(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`     // login via name
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var provider models.Provider
	query := `SELECT provider_id, name, password FROM providers WHERE name = $1`
	err = database.DB.QueryRow(query, input.Name).Scan(
		&provider.Provider_id, &provider.Name, &provider.Password,
	)

	if err != nil {
		log.Println("❌ Login error:", err)
		http.Error(w, "Provider not found", http.StatusUnauthorized)
		return
	}

	if provider.Password != input.Password {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Login successful",
		"provider_id": fmt.Sprintf("%d", provider.Provider_id),
		"name":        provider.Name,
	})
}

