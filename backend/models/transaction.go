package models

import "time"

type Transaction struct {
	TransactionID int       `json:"transaction_id"`
	RequestID     int       `json:"request_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Status        string    `json:"status"`
}
