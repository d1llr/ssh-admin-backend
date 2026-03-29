package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, c controllers.Controllers) {
	route := a.Group("/api/v1")

	route.Post("/user/sign/up", c.Auth.UserSignUp)
	route.Post("/user/sign/in", c.Auth.UserSignIn)
}
