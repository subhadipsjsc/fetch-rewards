package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the MySQL database
func InitDB() error {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return fmt.Errorf("environment variable DB_DSN not set")
	}

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Create table for receipts
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS receipts (
		id VARCHAR(255) PRIMARY KEY,
		points INT NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// SaveReceipt saves the receipt ID and points to the database
func SaveReceipt(id string, points int) error {
	_, err := db.Exec("INSERT INTO receipts (id, points) VALUES (?, ?)", id, points)
	return err
}

// GetPointsByID retrieves points by receipt ID
func GetPointsByID(id string) (int, error) {
	var points int

	err := db.QueryRow("SELECT points FROM receipts WHERE id = ?", id).Scan(&points)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("receipt not found")
	}
	return points, err
}
