package models

import (
	"time"

	"github.com/google/uuid"
)

type Secret struct {
	ID          uuid.UUID `json:"id"`
	Secret      string    `json:"secret"`
	Description string    `json:"description"`
	Timestamp   time.Time
}
