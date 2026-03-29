package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/consts"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case consts.AdminRoleName:
		// Nothing to do, verified successfully.
	case consts.ModeratorRoleName:
		// Nothing to do, verified successfully.
	case consts.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
