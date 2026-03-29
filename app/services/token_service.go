package services

import (
	"context"

	"github.com/google/uuid"
)

type Token interface {
	RenewToken(ctx context.Context, userId uuid.UUID) (*Tokens, error)
}
