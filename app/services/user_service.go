// app/services/auth_contract.go
package services

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/google/uuid"
)

type User interface {
	GetUserByID(id uuid.UUID) (models.User, error)
}
