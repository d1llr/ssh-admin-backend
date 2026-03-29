package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Check_exp(c *fiber.Ctx, claim TokenMetadata) error {
	now := time.Now().Unix()
	expires := claim.Expires

	if now > expires {
		return NewResponse(
			c,
			fiber.StatusUnauthorized,
			true,
			"Expired",
		)
	}
	return nil
}
