// File: backend/campaign_handlers.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"codex-arcana/backend/models"
)

// respondJSON envia um payload JSON com status customizado
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// sessionsHandler treats session endpoints for campaigns
func sessionsHandler(w http.ResponseWriter, r *http.Request) {
	// Treat GET /api/campaigns/{campaigID}/sessions
	if r.Method == http.MethodGet {
		// Extract campaign ID from the URL
		campaignID, err := extractCampaignIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid campaign ID", http.StatusBadRequest)
			return
		}
		// Retrieve sessions for the campaign
		sessions, err := GetSessionsByCampaign(campaignID)
		if err != nil {
			http.Error(w, "Failed to retrieve sessions", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, sessions)
		return
	}

	// Treat POST /api/campaigns/{campaignID}/sessions
	if r.Method == http.MethodPost {
		// Extract campaign ID from the URL
		campaignID, err := extractCampaignIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid campaign ID", http.StatusBadRequest)
			return
		}

		// Decodify payload into a Session model
		var s models.Session
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		// Validate non-empty title
		if strings.TrimSpace(s.Title) == "" {
			http.Error(w, "Session title cannot be empty", http.StatusBadRequest)
			return
		}

		// Set campaign ID and create the session
		s.CampaignID = campaignID
		created, err := CreateSession(s)
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusCreated, created)
		return
	}

	// Treat `PUT /api/campaigns/{campaignID}/sessions/{sessionID}`
	if r.Method == http.MethodPut {
		// Split the URL into segments
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 6 {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}
		// Extract IDs
		campaignID, err := extractCampaignIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid campaign ID", http.StatusBadRequest)
			return
		}

		sessionID, err := strconv.Atoi(parts[5])
		if err != nil {
			http.Error(w, "Invalid session ID", http.StatusBadRequest)
			return
		}

		// Decode payload into a Session model
		var s models.Session
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		// Validate non-empty title
		if strings.TrimSpace(s.Title) == "" {
			http.Error(w, "Session title cannot be empty", http.StatusBadRequest)
			return
		}

		// Adjust IDs and update the session
		s.ID = sessionID
		s.CampaignID = campaignID
		if err := UpdateSession(s); err != nil {
			http.Error(w, "Failed to update session", http.StatusInternalServerError)
			return
		}

		// Search and return the updated session
		updated, err := GetSessionByID(sessionID)
		if err != nil {
			http.Error(w, "Failed to retrieve updated session", http.StatusInternalServerError)
			return
		}

		respondJSON(w, http.StatusOK, updated)
		return
	}

	// Treat DELETE /api/campaigns/{campaignID}/sessions/{sessionID}
	if r.Method == http.MethodDelete {
		// Split the URL into segments
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 6 {
			http.Error(w, "Invalid session path", http.StatusBadRequest)
			return
		}
		sessionID, err := strconv.Atoi(parts[5])
		if err != nil {
			http.Error(w, "Invalid session ID", http.StatusBadRequest)
			return
		}

		if err := DeleteSession(sessionID); err != nil {
			http.Error(w, "Failed to delete session", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent) // No content to return on successful deletion
		return
	}
}

// parseCampaignRequest faz decode e valida nome de campanha
func parseCampaignRequest(r *http.Request) (models.Campaign, int, string) {
	var c models.Campaign
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return c, http.StatusBadRequest, "Invalid payload"
	}
	if strings.TrimSpace(c.Name) == "" {
		return c, http.StatusBadRequest, "Campaign name cannot be empty"
	}
	return c, 0, ""
}

// campaignsHandler treats GET /api/campaigns and POST /api/campaigns requests
func campaignsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// GET /api/campaigns
		camps, err := GetAllCampaigns()
		if err != nil {
			http.Error(w, "Failed to retrieve campaigns", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, camps)

	case http.MethodPost:
		// POST /api/campaigns
		c, status, msg := parseCampaignRequest(r)
		if msg != "" {
			http.Error(w, msg, status)
			return
		}
		created, err := CreateCampaign(c)
		if err != nil {
			http.Error(w, "Failed to create campaign", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusCreated, created)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// campaignHandler treats GET, PUT, and DELETE requests at /api/campaigns/{id}
func campaignHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for sessions
	if strings.Contains(r.URL.Path, "/sessions") {
		sessionsHandler(w, r)
		return
	}

	// Extract campaign ID from the URL
	id, err := extractCampaignIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid campaign ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// GET /api/campaigns/{id}
		c, err := GetCampaignByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Campaign not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Failed to retrieve campaign", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, c)

	case http.MethodPut:
		// PUT /api/campaigns/{id}
		c, status, msg := parseCampaignRequest(r)
		if msg != "" {
			http.Error(w, msg, status)
			return
		}
		c.ID = id
		if err := UpdateCampaign(c); err != nil {
			http.Error(w, "Failed to update campaign", http.StatusInternalServerError)
			return
		}
		updated, err := GetCampaignByID(id)
		if err != nil {
			http.Error(w, "Failed to retrieve updated campaign", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		// DELETE /api/campaigns/{id}
		if _, err := GetCampaignByID(id); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Campaign not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Failed to retrieve campaign", http.StatusInternalServerError)
			return
		}
		if err := DeleteCampaign(id); err != nil {
			http.Error(w, "Failed to delete campaign", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent) // No content to return on successful deletion

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// extractCampaignIDFromPath extracts the campaign ID from the URL path in the format /api/campaigns/{campaignID}/
func extractCampaignIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		return 0, fmt.Errorf("invalid path format")
	}
	return strconv.Atoi(parts[3])
}
