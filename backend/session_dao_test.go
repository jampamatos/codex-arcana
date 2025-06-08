// session_dao_test.go
package main

import (
	"database/sql"
	"testing"
	"time"

	"codex-arcana/backend/models"
)

func TestSessionDAO(t *testing.T) {
	// Inicializa o DB e garante que será fechado
	initDB()
	defer DB.Close()

	// Limpa tabelas para testar em ambiente limpo
	DB.Exec("DELETE FROM sessions")
	DB.Exec("DELETE FROM campaigns")

	// 1) Cria uma campanha para relacionar as sessões
	camp, err := CreateCampaign(models.Campaign{
		Name:        "Campanha para Teste",
		Description: "Descrição de campanha teste",
	})
	if err != nil {
		t.Fatalf("CreateCampaign falhou: %v", err)
	}
	if camp.ID == 0 {
		t.Fatalf("CreateCampaign retornou ID inválido")
	}

	// 2) Testa CreateSession
	now := time.Date(2025, 6, 10, 0, 0, 0, 0, time.UTC)
	sess, err := CreateSession(models.Session{
		CampaignID: camp.ID,
		Title:      "Sessão 1",
		Date:       now,
		Location:   "Taverna",
		Notes:      "Primeira sessão",
	})
	if err != nil {
		t.Fatalf("CreateSession falhou: %v", err)
	}
	if sess.ID == 0 {
		t.Fatalf("CreateSession retornou ID inválido")
	}

	// 3) Testa GetSessionsByCampaign
	list, err := GetSessionsByCampaign(camp.ID)
	if err != nil {
		t.Fatalf("GetSessionsByCampaign falhou: %v", err)
	}
	found := false
	for _, s := range list {
		if s.ID == sess.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("GetSessionsByCampaign não retornou a sessão criada")
	}

	// 4) Testa GetSessionByID
	fetched, err := GetSessionByID(sess.ID)
	if err != nil {
		t.Fatalf("GetSessionByID falhou: %v", err)
	}
	if fetched.Title != sess.Title || !fetched.Date.Equal(now) {
		t.Errorf("GetSessionByID retornou dados incorretos: %+v", fetched)
	}

	// 5) Testa UpdateSession
	fetched.Title = "Sessão Atualizada"
	fetched.Location = "Castelo"
	if err := UpdateSession(fetched); err != nil {
		t.Fatalf("UpdateSession falhou: %v", err)
	}
	updated, err := GetSessionByID(sess.ID)
	if err != nil {
		t.Fatalf("GetSessionByID após UpdateSession falhou: %v", err)
	}
	if updated.Title != "Sessão Atualizada" {
		t.Errorf("UpdateSession não atualizou o título: got %q", updated.Title)
	}

	// 6) Testa DeleteSession
	if err := DeleteSession(sess.ID); err != nil {
		t.Fatalf("DeleteSession falhou: %v", err)
	}
	_, err = GetSessionByID(sess.ID)
	if err == nil {
		t.Fatalf("Esperava erro ao buscar sessão deletada, mas não houve erro")
	}
	if err != sql.ErrNoRows {
		t.Fatalf("Esperava sql.ErrNoRows, mas obteve: %v", err)
	}
}
