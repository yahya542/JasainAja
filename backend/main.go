package main

import (
	"log"
	"net/http"

	"jasainaja-backend/database"
	"jasainaja-backend/handlers"
	"github.com/gorilla/mux"
)

// CORS Middleware (global for all routes)
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Connect()

	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")

	// Provider routes
	router.HandleFunc("/api/provider/register", handlers.RegisterProvider).Methods("POST")
	router.HandleFunc("/api/provider/login", handlers.LoginProvider).Methods("POST")

	// Service routes
	router.HandleFunc("/api/services", handlers.CreateService).Methods("POST")
	router.HandleFunc("/api/services/all", handlers.GetAllServices).Methods("GET")
	router.HandleFunc("/api/services/provider/{id}", handlers.GetServicesByProviderID).Methods("GET")

	// Request & Transaction routes
	router.HandleFunc("/api/request", handlers.CreateRequest).Methods("POST")
	router.HandleFunc("/api/requests/user/{id}", handlers.GetRequestsByUserID).Methods("GET")
	router.HandleFunc("/api/requests/provider/{id}", handlers.GetRequestsByProviderID).Methods("GET")
	router.HandleFunc("/api/request/confirm", handlers.ConfirmRequest).Methods("POST")
	router.HandleFunc("/api/transactions", handlers.GetTransactionsByRequestID).Methods("GET")
	router.HandleFunc("/api/transactions/complete", handlers.CompleteTransaction).Methods("POST")

	// Admin routes
	router.HandleFunc("/api/admin/reports", handlers.GetAdminReports).Methods("GET")
	router.HandleFunc("/api/admin/reports/mark-read", handlers.MarkReportRead).Methods("POST")

	// Wrap router with CORS middleware
	log.Println("✅ Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(router)))
}
