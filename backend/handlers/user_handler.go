package handlers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "jasainaja-backend/database"
    "jasainaja-backend/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (username, email, password, time_balance) VALUES ($1, $2, $3, $4)"
	_, err = database.DB.Exec(query, user.Name, user.Email, user.Password, user.TimeBalance)

	if err != nil {
		log.Println("❌ Error insert user:", err) // <--- tambahkan ini
		http.Error(w, "Failed to register", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT user_id, username, password FROM users WHERE username = $1 AND password = $2`
	err = database.DB.QueryRow(query, input.Username, input.Password).Scan(
		&user.UserID, &user.Name, &user.Password,
	)

	if err != nil {
		log.Println("❌ Login error:", err)
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"user_id": fmt.Sprintf("%d", user.UserID),
		"username": user.Name,
	})
}


