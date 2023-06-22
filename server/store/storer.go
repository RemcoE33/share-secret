package store

import (
	"github.com/RemcoE33/share-secret/models"
	"github.com/google/uuid"
)

type Storer interface {
	SaveSecret(secret *models.Secret) error
	GetSecret(id uuid.UUID) (models.Secret, error)
	DeleteSecret(secret *models.Secret) error
	DeleteAllPasedTTLSecrets() error
}
