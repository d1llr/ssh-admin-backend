// app/repositories/user_repository.go
package repositories

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/queries"
	"github.com/google/uuid"
)

type SshRepository interface {
	GetAllSSHConnectionsByUserId(id uuid.UUID) ([]models.SSH, error)
	CreateSSHConnection(s *models.SSH) error
}

type sshRepository struct {
	q *queries.SSHQueries
}

func NewSshRepository(q *queries.SSHQueries) SshRepository {
	return &sshRepository{q: q}
}

func (r *sshRepository) GetAllSSHConnectionsByUserId(id uuid.UUID) ([]models.SSH, error) {
	return r.q.GetAllSSHConnectionsByUserId(id)
}

func (r *sshRepository) CreateSSHConnection(s *models.SSH) error {
	return r.q.CreateSSHConnection(s)
}
