// File: backend/session_dao.go

package main

import (
	"database/sql"
	"time"

	"codex-arcana/backend/models"
)

// CreateSession inserts a new session linked to a campaign and returns the created registry.
func CreateSession(s models.Session) (models.Session, error) {
	// Define timestamps
	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now

	// Execute INSERT and get the generated ID
	result, err := DB.Exec(
		`INSERT INTO sessions
		(campaign_id, title, date, location, notes, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		s.CampaignID,
		s.Title,
		s.Date,
		s.Location,
		s.Notes,
		s.CreatedAt,
		s.UpdatedAt,
	)
	if err != nil {
		return models.Session{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Session{}, err
	}

	s.ID = int(id)

	return s, nil
}

// GetSessionsByCampaign returns all sessions linked to a specific campaign.
func GetSessionsByCampaign(campaignID int) ([]models.Session, error) {
	var sessions []models.Session

	rows, err := DB.Query(
		`SELECT id, campaign_id, title, date, location, notes, created_at, updated_at
		FROM sessions
		WHERE campaign_id = ?`,
		campaignID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Session
		if err := rows.Scan(
			&s.ID,
			&s.CampaignID,
			&s.Title,
			&s.Date,
			&s.Location,
			&s.Notes,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

// GetSessionByID returns a session by its ID.
func GetSessionByID(id int) (models.Session, error) {
	var s models.Session
	row := DB.QueryRow(
		`SELECT id, campaign_id, title, date, location, notes, created_at, updated_at
		FROM sessions
		WHERE id = ?`,
		id,
	)

	if err := row.Scan(
		&s.ID,
		&s.CampaignID,
		&s.Title,
		&s.Date,
		&s.Location,
		&s.Notes,
		&s.CreatedAt,
		&s.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Session{}, err
		}
		return models.Session{}, err
	}
	return s, nil
}

// UpdateSession updatres an existing session
func UpdateSession(s models.Session) error {
	// Update timestamps
	s.UpdatedAt = time.Now()

	// Execute UPDATE statement
	_, err := DB.Exec(
		`UPDATE sessions
		SET title = ?, date = ?, location = ?, notes = ?, updated_at = ?
		WHERE id = ?`,
		s.Title,
		s.Date,
		s.Location,
		s.Notes,
		s.UpdatedAt,
		s.ID,
	)

	return err
}

// DeleteSession removes a session by its ID.
func DeleteSession(id int) error {
	_, err := DB.Exec(
		`DELETE FROM sessions
		WHERE id = ?`,
		id,
	)
	return err
}
