package store

import (
	"context"

	"github.com/RemcoE33/share-secret/models"
	"github.com/google/uuid"
)

type Store interface {
	SaveSecret(ctx context.Context, secret *models.Secret) (uuid.UUID, error)
	GetSecret(ctx context.Context, id uuid.UUID) (models.Secret, error)
	DeleteSecret(ctx context.Context, secret *models.Secret) error
	DeleteAllPasedTTLSecrets() error
}
