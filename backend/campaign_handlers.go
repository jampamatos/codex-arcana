package main

import (
	"database/sql"
	"encoding/json"
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
	// Extract campaign ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/campaigns/")
	id, err := strconv.Atoi(idStr)
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
