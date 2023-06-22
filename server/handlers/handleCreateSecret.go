package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RemcoE33/share-secret/models"
	"github.com/google/uuid"
)

// all the logic for creating the secret is in the handler
// calculates the ttl, encrypt the secret and generates the uuid
// returns the models.IdResponse to the client
func (s *Server) HandleCreateSecret(w http.ResponseWriter, r *http.Request) {
	var secret models.Secret

	if err := json.NewDecoder(r.Body).Decode(&secret); err != nil {
		serveInternalError(err, w)
		return
	}

	//Set TTL
	now := time.Now()
	dur := time.Duration(secret.Days)
	ttl := now.Add(dur * (time.Hour * 24))
	secret.ExpiresAt = ttl

	//Creating uuid for url
	id, err := uuid.NewUUID()
	if err != nil {
		serveInternalError(err, w)
		return
	}

	secret.ID = id.String()

	//Save in store
	err = s.Storer.SaveSecret(&secret)
	if err != nil {
		serveInternalError(err, w)
		return
	}

	serveJSON(secret.ToIdResponse(), w)

}
