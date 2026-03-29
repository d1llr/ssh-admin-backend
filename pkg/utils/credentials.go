package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/consts"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case consts.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
			consts.BookCreateCredential,
			consts.BookUpdateCredential,
			consts.BookDeleteCredential,
		}
	case consts.ModeratorRoleName:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			consts.BookCreateCredential,
			consts.BookUpdateCredential,
		}
	case consts.UserRoleName:
		// Simple user credentials (only book creation).
		credentials = []string{
			consts.BookCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
