package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateSSH struct to create new ssh connection
type SSH struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UserID       uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Host         string    `db:"email" json:"host" validate:"required,lte=255"`
	Username     string    `db:"username" json:"username" validate:"required,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
}
