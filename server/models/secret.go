package models

import (
	"time"
)

type Secret struct {
	IdResponse
	GetSecretResponse
	Days      int64     `json:"days" firestore:"days"`
	ExpiresAt time.Time `firestore:"expiresAt"`
}

// Response to the client when there is a secret created
type IdResponse struct {
	ID string `json:"id" firestore:"id"`
}

// When a secret is requested from the client
type GetSecretResponse struct {
	Secret string `json:"secret" firestore:"secret"`
}

// convert the full secret stuct to a response for the client
func (s *Secret) ToSecretResponse() GetSecretResponse {
	return GetSecretResponse{
		Secret: s.Secret,
	}
}

// conver the full secret struct to a response for the client
func (s *Secret) ToIdResponse() IdResponse {
	return IdResponse{
		ID: s.ID,
	}
}
