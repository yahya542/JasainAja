package main 

import (
	"log"
	"net/http"

	"jasainaja-backend/database"
	"jasainaja-backend/handlers"

)


func main() {

	database.Connect()

	http.HandleFunc("/api/register", handlers.RegisterUser)
	http.HandleFunc("/api/login", handlers.LoginUser)
	http.HandleFunc("/api/provider/register", handlers.RegisterProvider)
	http.HandleFunc("/api/provider/login", handlers.LoginProvider)
    http.HandleFunc("/api/services", handlers.CreateService)              // POST
    http.HandleFunc("/api/services/all", handlers.GetAllServices)        // GET
    http.HandleFunc("/api/services/provider/", handlers.GetServicesByProviderID) // GET
    http.HandleFunc("/api/request", handlers.CreateRequest)                        // POST
	http.HandleFunc("/api/requests/user/", handlers.GetRequestsByUserID)          // GET
	http.HandleFunc("/api/requests/provider/", handlers.GetRequestsByProviderID)  // GET
	http.HandleFunc("/api/request/confirm", handlers.ConfirmRequest)              // POST
    http.HandleFunc("/api/transactions", handlers.GetTransactionsByRequestID) // POST
	http.HandleFunc("/api/transactions/complete", handlers.CompleteTransaction) // POST
	http.HandleFunc("/api/admin/reports", handlers.GetAdminReports)       // GET
	http.HandleFunc("/api/admin/reports/mark-read", handlers.MarkReportRead) // POST



	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
