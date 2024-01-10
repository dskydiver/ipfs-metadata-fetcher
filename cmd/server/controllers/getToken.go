package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dskydiver/ipfs-metadata-fetcher/ent"
	"github.com/dskydiver/ipfs-metadata-fetcher/ent/metadata"
	"github.com/gorilla/mux"
)

func (tc *TokenController) GetToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	cid, _ := vars["cid"]

	token, err := tc.client.Metadata.Query().Where(
		metadata.CidEQ(cid),
	).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "token not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, fmt.Sprintf("error getting token: %s", err), http.StatusInternalServerError)
			return
		}
	}

	tokenRes := TokenResult{
		ID:          token.ID,
		CID:         token.Cid,
		Image:       token.Image,
		Description: token.Description,
		Name:        token.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenRes)
}
