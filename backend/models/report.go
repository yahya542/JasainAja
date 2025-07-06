package models

import "time"

type AdminReport struct {
	ReportID      int       `json:"report_id"`
	TransactionID int       `json:"transaction_id"`
	ReportType    string    `json:"report_type"`
	CreatedAt     time.Time `json:"created_at"`
	Status        string    `json:"status"`
}
