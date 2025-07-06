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


	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
