package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"jasainaja-backend/database"
	"jasainaja-backend/models"
)

// Buat transaksi otomatis ketika request dikonfirmasi
func CreateTransactionFromRequest(requestID int) error {
	query := `INSERT INTO transactions (request_id, start_time, status) VALUES ($1, $2, 'in_progress')`
	_, err := database.DB.Exec(query, requestID, time.Now())
	return err
}

// Ambil transaksi by user
func GetTransactionsByRequestID(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		RequestID int `json:"request_id"`
	}
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	row := database.DB.QueryRow(`SELECT transaction_id, request_id, start_time, end_time, status FROM transactions WHERE request_id = $1`, req.RequestID)

	var tx models.Transaction
	err := row.Scan(&tx.TransactionID, &tx.RequestID, &tx.StartTime, &tx.EndTime, &tx.Status)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tx)
}

// Ubah status transaksi jadi completed
func CompleteTransaction(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		TransactionID int `json:"transaction_id"`
	}
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	_, err := database.DB.Exec(`UPDATE transactions SET status = 'completed', end_time = $1 WHERE transaction_id = $2`, time.Now(), req.TransactionID)
	if err != nil {
		http.Error(w, "Failed to complete transaction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Transaction completed"})
}
