// File: backend/database.go

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Global variable to hold the database connection
var DB *sql.DB

// initDB opens a connection to the SQLite database and initializes the Campaign table, if it doesn't exist.
func initDB() {
	var err error
	// Open (or create) the SQLite database file
	DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Optional: configure timeouts/pragmas, but we'll use defaults for simplicity

	// SQL statement to create the Campaign table if it doesn't exist
	createCampaignTable := `
	CREATE TABLE IF NOT EXISTS campaigns (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
		);
	`

	if _, err = DB.Exec(createCampaignTable); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createSessionsTable := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		campaign_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		date DATETIME NOT NULL,
		location TEXT,
		notes TEXT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE
		)`

	if _, err = DB.Exec(createSessionsTable); err != nil {
		log.Fatalf("Failed to create sessions table: %v", err)
	}
}
