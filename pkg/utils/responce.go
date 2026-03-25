package utils

import "github.com/gofiber/fiber/v2"

func NewResponse(c *fiber.Ctx, status int, err bool, msg string, data ...fiber.Map) error {
	response := fiber.Map{
		"error": err,
		"msg":   msg,
	}

	// мержим все переданные поля
	for _, d := range data {
		for k, v := range d {
			response[k] = v
		}
	}

	return c.Status(status).JSON(response)
}
