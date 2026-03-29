package services

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/models"
)

type Tokens struct {
	Access  string
	Refresh string
}

type Auth interface {
	SignUp(signUp *models.SignUp) (*models.User, error)
	SignIn(ctx context.Context, in *models.SignIn) (*Tokens, error)
}
