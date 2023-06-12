package store

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/RemcoE33/share-secret/models"
	"github.com/google/uuid"
)

var (
	firestoreProjectId = "investservices"
	collection         = "secrets"
)

type Client firestore.Client

type Firestore struct {
	client     *Client
	projectId  string
	collection string
}

func (f *Firestore) SaveSecret(ctx context.Context, s *models.Secret) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id, fmt.Errorf("could not create an uuid")
	}

	_, err = f.client.Collection(f.collection).Doc(id).Set(ctx, s)
	if err != nil {
		return id, fmt.Errorf("could not save to store")
	}

	return id, err
}

// Gets the secret from the store
func (f *Firestore) GetSecret(ctx context.Context, id uuid.UUID) (*models.Secret, error) {
	var s models.Secret

	snap, err := f.client.Collection(f.collection).Doc(id).Get(ctx)
	if err != nil {
		return &s, fmt.Errorf("could not get secret: %v", err)
	}

	snap.DataTo(&s)
	return &s, nil

}

// Deleting a secret after it is returned to the client
func (f *Firestore) DeleteSecret(ctx context.Context, s *models.Secret) error {
	_, err := f.client.Collection(f.collection).Doc(s.ID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("could not delete %v from store", s.ID)
	}
	return nil
}

// Gets all the secrets older then the limited time and deletes it.
func (f *Firestore) DeleteAllPasedTTLSecrets(t time.Time) error {
	return nil
}

// Initialize a new store
func NewFireBaseStore(ctx context.Context) (*Firestore, error) {
	client, err := firestore.NewClient(ctx, firestoreProjectId)
	if err != nil {
		return nil, fmt.Errorf("unable to make connection to firestore")
	}

	return &Firestore{
		client:     client,
		collection: collection,
	}, nil
}
