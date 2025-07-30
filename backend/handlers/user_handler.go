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

    query := `INSERT INTO users (username, email, password, user_type) VALUES ($1, $2, $3, $4)`
    _, err = database.DB.Exec(query, user.Username, user.Email, user.Password, "client")
    if err != nil {
        log.Println("❌ Error insert user:", err)
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
    query := `SELECT user_id, username, password FROM users WHERE username = $1 AND user_type = 'client'`
    err = database.DB.QueryRow(query, input.Username).Scan(
        &user.UserID, &user.Username, &user.Password,
    )
    if err != nil {
        log.Println("❌ Login error:", err)
        http.Error(w, "Username not found", http.StatusUnauthorized)
        return
    }

    if user.Password != input.Password {
        http.Error(w, "Incorrect password", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message":  "Login successful",
        "user_id":  fmt.Sprintf("%d", user.UserID),
        "username": user.Username,
    })
}
