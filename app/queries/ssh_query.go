package queries

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/jmoiron/sqlx"
)

type SSHQueries struct {
	*sqlx.DB
}

func (q *BookQueries) GetAllSSHConnections() ([]models.SSH, error) {
	// Define books variable.
	ssh := []models.SSH{}

	// Define query string.
	query := `SELECT * FROM ssh_connections`

	// Send query to database.
	err := q.Select(&ssh, query)
	if err != nil {
		// Return empty object and error.
		return ssh, err
	}

	// Return query result.
	return ssh, nil
}

func (q *BookQueries) CreateSSHConnection(b *models.SSH) error {
	// Define query string.
	query := `INSERT INTO ssh_connections VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Host, b.PasswordHash, b.Username)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
