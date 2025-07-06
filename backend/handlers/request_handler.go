package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"jasainaja-backend/database"
	"jasainaja-backend/models"
)

// 1. Buat request jasa
func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO request (user_id, provider_id, description, status) VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(query, req.UserID, req.ProviderID, req.Description, "pending")
	if err != nil {
		log.Println("❌ Failed to create request:", err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Request created successfully"})
}

// 2. Ambil semua request milik user
func GetRequestsByUserID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[4])

	rows, err := database.DB.Query(`SELECT request_id, user_id, provider_id, description, status, created_at FROM request WHERE user_id = $1`, id)
	if err != nil {
		http.Error(w, "Error fetching user requests", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []models.Request
	for rows.Next() {
		var req models.Request
		rows.Scan(&req.RequestID, &req.UserID, &req.ProviderID, &req.Description, &req.Status, &req.CreatedAt)
		requests = append(requests, req)
	}

	json.NewEncoder(w).Encode(requests)
}

// 3. Ambil semua request milik provider
func GetRequestsByProviderID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[4])

	rows, err := database.DB.Query(`SELECT request_id, user_id, provider_id, description, status, created_at FROM request WHERE provider_id = $1`, id)
	if err != nil {
		http.Error(w, "Error fetching provider requests", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []models.Request
	for rows.Next() {
		var req models.Request
		rows.Scan(&req.RequestID, &req.UserID, &req.ProviderID, &req.Description, &req.Status, &req.CreatedAt)
		requests = append(requests, req)
	}

	json.NewEncoder(w).Encode(requests)
}

// 4. Konfirmasi request → ubah status jadi confirmed
func ConfirmRequest(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RequestID int `json:"request_id"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	_, err := database.DB.Exec(`UPDATE request SET status = 'confirmed' WHERE request_id = $1`, input.RequestID)
	if err != nil {
		http.Error(w, "Failed to confirm request", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Request confirmed"})
	_ = CreateTransactionFromRequest(input.RequestID)

}
