package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// get the secret from store and return to the client
// get id from chi url param, checks ttl, decrypt secret
// returns models.SecretResponse to the client
func (s *Server) HandleGetSecret(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		serveNotFoundOrPastDeadline(w)
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		serveNotFoundOrPastDeadline(w)
		return
	}

	secret, err := s.Storer.GetSecret(uuid)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			serveNotFoundOrPastDeadline(w)
		} else {
			serveInternalError(err, w)
		}
		return
	}

	if secret.Secret == "" {
		serveNotFoundOrPastDeadline(w)
		return
	}

	now := time.Now()
	ex := secret.ExpiresAt
	if now.After(ex) {
		s.Storer.DeleteSecret(&secret)
		serveNotFoundOrPastDeadline(w)
		return
	}

	if err = serveJSON(secret.ToSecretResponse(), w); err == nil {
		if err := s.Storer.DeleteSecret(&secret); err != nil {
			msg := fmt.Errorf("could not delete id: %s -> %v", id, err)
			fmt.Println(msg)
		}
	}
}
