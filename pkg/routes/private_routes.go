package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes Защищенные методы AuthGuard
func PrivateRoutes(a *fiber.App, c controllers.Controllers) {
	route := a.Group("/api/v1", middleware.AuthGuard())

	route.Post("/ssh/create", c.Ssh.CreateSSH)
	route.Get("/ssh/all", c.Ssh.GetAllSSHConnections)

	route.Post("/token/renew", c.Token.RenewTokens)
}
