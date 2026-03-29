package httpctx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/create-go-app/fiber-go-template/pkg/utils"
)

const (
	keyUserID      = "userID"
	keyCredentials = "credentials"
	keyTokenMeta   = "tokenMeta"
)

// Setters

func SetUserID(c *fiber.Ctx, id uuid.UUID) {
	c.Locals(keyUserID, id)
}

func SetCredentials(c *fiber.Ctx, creds map[string]bool) {
	c.Locals(keyCredentials, creds)
}

func SetTokenMeta(c *fiber.Ctx, meta *utils.TokenMetadata) {
	c.Locals(keyTokenMeta, meta)
}

func UserID(c *fiber.Ctx) uuid.UUID {
	v := c.Locals(keyUserID)
	id := v.(uuid.UUID)
	return id
}

func Credentials(c *fiber.Ctx) map[string]bool {
	v := c.Locals(keyCredentials)
	creds := v.(map[string]bool)
	return creds
}

func TokenMeta(c *fiber.Ctx) *utils.TokenMetadata {
	v := c.Locals(keyTokenMeta)
	meta := v.(*utils.TokenMetadata)
	return meta
}
