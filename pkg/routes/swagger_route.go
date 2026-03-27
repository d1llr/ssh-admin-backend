package routes

import (
	"github.com/create-go-app/fiber-go-template/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(app *fiber.App) {

	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendString(docs.SwaggerInfo.ReadDoc())
	})
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/debug/doc.json", func(c *fiber.Ctx) error {
		return c.SendString(docs.SwaggerInfo.ReadDoc())
	})
}
