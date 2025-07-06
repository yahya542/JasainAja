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

// Tambah jasa
func CreateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO services (provider_id, title, description, duration_minutes, price_time) VALUES ($1, $2, $3, $4, $5)`
	_, err = database.DB.Exec(query, service.ProviderID, service.Title, service.Description, service.DurationMinutes, service.PriceTime)
	if err != nil {
		log.Println("❌ Failed to insert service:", err)
		http.Error(w, "Failed to create service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Service created successfully"})
}

func GetAllServices(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT service_id, provider_id, title, description, duration_minutes, price_time FROM services`)
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ServiceID, &s.ProviderID, &s.Title, &s.Description, &s.DurationMinutes, &s.PriceTime)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		services = append(services, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}
func GetServicesByProviderID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	providerID, err := strconv.Atoi(parts[4])
	if err != nil {
		http.Error(w, "Invalid provider ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`SELECT service_id, provider_id, title, description, duration_minutes, price_time FROM services WHERE provider_id = $1`, providerID)
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ServiceID, &s.ProviderID, &s.Title, &s.Description, &s.DurationMinutes, &s.PriceTime)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		services = append(services, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}