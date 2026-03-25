package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Check_exp(c *fiber.Ctx, claim TokenMetadata) error {
	now := time.Now().Unix()
	expires := claim.Expires

	if now > expires {
		return NewResponse(
			c,
			fiber.StatusUnauthorized,
			true,
			"unauthorized, check expiration time of your token",
		)
	}
	return nil
}
