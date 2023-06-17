package link

import (
	"time"

	"github.com/google/uuid"
)

type (
	linkCreate struct {
		Owner    uuid.UUID `json:"-"`
		Original string    `json:"original" validate:"required"`
		Trimmed  string    `json:"trimmed" validate:"required"`
	}

	linkUpdate struct {
		Original string `json:"original" validate:"required"`
		Trimmed  string `json:"trimmed" validate:"required"`
	}

	Link struct {
		ID        uuid.UUID `json:"id"`
		OwnerID   uuid.UUID `json:"owner_id"`
		Original  string    `json:"original"`
		Trimmed   string    `json:"trimmed"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
