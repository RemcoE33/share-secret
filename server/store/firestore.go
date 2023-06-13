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
	Client     *firestore.Client
	ProjectId  string
	Collection *firestore.CollectionRef
}

// Initialize a new store
func NewFirestore(project, collection string) (*Firestore, error) {
	client, err := firestore.NewClient(context.Background(), firestoreProjectId)
	if err != nil {
		return nil, fmt.Errorf("unable to make connection to firestore")
	}

	col := client.Collection(collection)

	return &Firestore{
		Client:     client,
		Collection: col,
		ProjectId:  project,
	}, nil
}

func (f *Firestore) SaveSecret(s *models.Secret) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id, fmt.Errorf("could not create an uuid")
	}

	_, err = f.client.Collection(f.collection).Doc(id).Set(context.Background(), s)
	if err != nil {
		return id, fmt.Errorf("could not save to store")
	}

	return id, err
}

// Gets the secret from the store
func (f *Firestore) GetSecret(id uuid.UUID) (*models.Secret, error) {
	var s models.Secret

	snap, err := f.client.Collection(f.collection).Doc(id).Get(context.Background())
	if err != nil {
		return &s, fmt.Errorf("could not get secret: %v", err)
	}

	snap.DataTo(&s)
	return &s, nil

}

// Deleting a secret after it is returned to the client
func (f *Firestore) DeleteSecret(s *models.Secret) error {
	_, err := f.client.Collection(f.collection).Doc(s.ID).Delete(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete %v from store", s.ID)
	}
	return nil
}

// Gets all the secrets older then the limited time and deletes it.
func (f *Firestore) DeleteAllPasedTTLSecrets(t time.Time) error {
	return nil
}
