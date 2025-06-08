// File: backend/models/models.go

package models

import "time"

// Campaign represents an RPG campaign in the database.
type Campaign struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Session represents a session within a Campaign.
type Session struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	Location   string    `json:"location"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
