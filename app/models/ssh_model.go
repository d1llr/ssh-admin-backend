package models

import (
	"time"

	"github.com/google/uuid"
)

// SSH struct to describe SSH object.
type SSH struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Host      string    `db:"host" json:"host" validate:"required,lte=255"`
	Name      string    `db:"name" json:"name" validate:"required,lte=255"`
	Password  string    `db:"password" json:"password,omitempty" validate:"required,lte=255"`
}
