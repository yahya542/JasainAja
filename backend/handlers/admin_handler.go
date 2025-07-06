package handlers

import (
	"encoding/json"
	"net/http"

	"jasainaja-backend/database"
	"jasainaja-backend/models"
)

// Lihat semua laporan admin
func GetAdminReports(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT report_id, transaction_id, report_type, created_at, status FROM admin_reports`)
	if err != nil {
		http.Error(w, "Failed to fetch reports", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reports []models.AdminReport
	for rows.Next() {
		var r models.AdminReport
		rows.Scan(&r.ReportID, &r.TransactionID, &r.ReportType, &r.CreatedAt, &r.Status)
		reports = append(reports, r)
	}

	json.NewEncoder(w).Encode(reports)
}

// Tandai laporan sudah dibaca
func MarkReportRead(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ReportID int `json:"report_id"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	_, err := database.DB.Exec(`UPDATE admin_reports SET status = 'read' WHERE report_id = $1`, input.ReportID)
	if err != nil {
		http.Error(w, "Failed to update report", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Report marked as read"})
}
