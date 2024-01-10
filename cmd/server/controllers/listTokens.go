package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dskydiver/ipfs-metadata-fetcher/ent"
)

type TokenController struct {
	client *ent.Client
}

type TokenResult struct {
	ID          string `json:"id"`
	CID         string `json:"cid"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

func NewTokenController(client *ent.Client) *TokenController {
	return &TokenController{
		client: client,
	}
}

func (tc *TokenController) ListTokens(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	tokens, err := tc.client.Metadata.Query().All(ctx)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error getting tokens: %s", err), http.StatusInternalServerError)
		return
	}

	tokensRes := make([]TokenResult, len(tokens))

	for i, token := range tokens {
		tokensRes[i] = TokenResult{
			ID:          token.ID,
			CID:         token.Cid,
			Image:       token.Image,
			Description: token.Description,
			Name:        token.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokensRes)
}
