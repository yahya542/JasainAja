package main

import (
	"log"
	"net/http"

	"jasainaja-backend/database"
	"jasainaja-backend/handlers"
)

// CORS middleware untuk semua endpoint
func withCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r)
	}
}

func main() {
	database.Connect()

	http.HandleFunc("/api/register", withCORS(handlers.RegisterUser))
	http.HandleFunc("/api/login", withCORS(handlers.LoginUser))
	http.HandleFunc("/api/provider/register", withCORS(handlers.RegisterProvider))
	http.HandleFunc("/api/provider/login", withCORS(handlers.LoginProvider))
	http.HandleFunc("/api/services", withCORS(handlers.CreateService))              
	http.HandleFunc("/api/services/all", withCORS(handlers.GetAllServices))        
	http.HandleFunc("/api/services/provider/", withCORS(handlers.GetServicesByProviderID)) 
	http.HandleFunc("/api/request", withCORS(handlers.CreateRequest))                        
	http.HandleFunc("/api/requests/user/", withCORS(handlers.GetRequestsByUserID))          
	http.HandleFunc("/api/requests/provider/", withCORS(handlers.GetRequestsByProviderID))  
	http.HandleFunc("/api/request/confirm", withCORS(handlers.ConfirmRequest))              
	http.HandleFunc("/api/transactions", withCORS(handlers.GetTransactionsByRequestID)) 
	http.HandleFunc("/api/transactions/complete", withCORS(handlers.CompleteTransaction)) 
	http.HandleFunc("/api/admin/reports", withCORS(handlers.GetAdminReports))       
	http.HandleFunc("/api/admin/reports/mark-read", withCORS(handlers.MarkReportRead)) 

	log.Println("✅ Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
