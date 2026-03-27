package queries

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SSHQueries struct {
	*sqlx.DB
}

func (q *SSHQueries) GetAllSSHConnectionsByUserId(userId uuid.UUID) ([]models.SSH, error) {
	ssh := []models.SSH{}

	query := `
		SELECT id, host, name, password, created_at, user_id
		FROM ssh_connections
		where user_id=$1
	`
	err := q.Select(&ssh, query, userId)
	if err != nil {
		// Return empty object and error.
		return ssh, err
	}

	return ssh, nil
}

func (q *SSHQueries) CreateSSHConnection(b *models.SSH) error {
	query := `
		INSERT INTO ssh_connections
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.Exec(query, b.ID, b.Name, b.Host, b.UserID, b.Password, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}
