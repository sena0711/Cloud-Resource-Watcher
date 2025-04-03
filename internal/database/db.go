package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "host=postgres user=admin password=password dbname=aws_asset_tracker sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected successfully!")
}
