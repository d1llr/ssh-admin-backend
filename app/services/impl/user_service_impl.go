// app/services/user_service.go
package impl

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repositories"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/google/uuid"
)

type UserService struct {
	repos repositories.Repositories
}

func NewUserService(repos repositories.Repositories) services.User {
	return &UserService{repos: repos}
}

func (s *UserService) GetUserByID(id uuid.UUID) (models.User, error) {
	return s.repos.User.GetByID(id)
}
