// app/services/auth_contract.go
package services

import (
	"github.com/gofiber/fiber/v2"
)

type Ssh interface {
	CreateSSHConnection(c *fiber.Ctx) error
	GetAllSSHConnectionsByUserID(c *fiber.Ctx) error
}
