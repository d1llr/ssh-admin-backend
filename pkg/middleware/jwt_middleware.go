package middleware

import (
	"time"

	"github.com/create-go-app/fiber-go-template/pkg/httpctx"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		meta, err := utils.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "invalid or missing token",
			})
		}

		now := time.Now().Unix()
		if now > meta.Expires {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "token expired",
			})
		}

		httpctx.SetUserID(c, meta.UserID)
		httpctx.SetCredentials(c, meta.Credentials)
		httpctx.SetTokenMeta(c, meta)

		return c.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
