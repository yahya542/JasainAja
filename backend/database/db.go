package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:deya2501@127.0.0.1:5432/jasainaja_db?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}
