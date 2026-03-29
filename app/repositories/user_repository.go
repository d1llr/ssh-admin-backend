// app/repositories/user_repository.go
package repositories

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/queries"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetByID(id uuid.UUID) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Create(u *models.User) error
}

type userRepository struct {
	q *queries.UserQueries
}

func NewUserRepository(q *queries.UserQueries) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) GetByID(id uuid.UUID) (models.User, error) {
	return r.q.GetUserByID(id)
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	return r.q.GetUserByEmail(email)
}

func (r *userRepository) Create(u *models.User) error {
	return r.q.CreateUser(u)
}
