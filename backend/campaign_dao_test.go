package main

import (
	"database/sql"
	"testing"

	"codex-arcana/backend/models"
)

// TestCampaignDAO executes a series o CRUD operations to check the Campaign DAO functionality
func TestCampaignDAO(t *testing.T) {
	initDB()         // Initialize the database connection
	defer DB.Close() // Ensure the database is closed after tests

	// Create a new campaign
	c := models.Campaign{
		Name:        "Test Campaign",
		Description: "This is a test campaign",
	}
	c, err := CreateCampaign(c)
	if err != nil {
		t.Fatalf("CreateCampaign failed to create campaign: %v", err)
	}
	// Check if the campaign was created correctly
	if c.ID == 0 {
		t.Fatalf("Expected campaign ID to be set, got %d", c.ID)
	}
	// Check if the timestamps are set correctly
	if c.CreatedAt.IsZero() || c.UpdatedAt.IsZero() {
		t.Fatalf("Expected timestamps to be set, got CreatedAt: %v, UpdatedAt: %v", c.CreatedAt, c.UpdatedAt)
	}
	// Check if GetAllCampaigns returns at least the created campaign
	campaigns, err := GetAllCampaigns()
	if err != nil {
		t.Fatalf("GetAllCampaigns failed: %v", err)
	}
	if len(campaigns) == 0 {
		t.Fatalf("Expected at least one campaign, got none")
	}

	found := false
	for _, campaign := range campaigns {
		if campaign.ID == c.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Created campaign not found in GetAllCampaigns result")
	}

	// Check GetCampaignByID
	campaign, err := GetCampaignByID(c.ID)
	if err != nil {
		t.Fatalf("GetCampaignByID failed: %v", err)
	}
	if campaign.Name != c.Name || campaign.Description != c.Description {
		t.Fatalf("GetCampaignByID returned incorrect campaign: got %v, want %v", campaign, c)
	}

	// Check UpdateCampaign
	c.Name = "Updated Campaign"
	c.Description = "This is an updated test campaign"
	err = UpdateCampaign(c)
	if err != nil {
		t.Fatalf("UpdateCampaign failed: %v", err)
	}
	// Check campaign ID again
	c, err = GetCampaignByID(c.ID)
	if err != nil {
		t.Fatalf("GetCampaignByID after update failed: %v", err)
	}
	if c.Name != "Updated Campaign" || c.Description != "This is an updated test campaign" {
		t.Fatalf("UpdateCampaign did not update campaign correctly: got %v, want %v", c, "Updated Campaign")
	}

	// Check DeleteCampaign
	err = DeleteCampaign(c.ID)
	if err != nil {
		t.Fatalf("DeleteCampaign failed: %v", err)
	}
	// Check if the campaign is deleted
	_, err = GetCampaignByID(c.ID)
	if err == nil {
		t.Fatalf("Expected GetCampaignByID to return an error for deleted campaign, got none")
	}
	if err != sql.ErrNoRows {
		t.Fatalf("Expected GetCampaignByID to return sql.ErrNoRows, got %v", err)
	}

	// Clean up table after tests
	_, _ = DB.Exec("DELETE FROM campaigns")
}
