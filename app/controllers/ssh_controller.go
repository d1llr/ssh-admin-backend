package controllers

import (
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/gofiber/fiber/v2"
)

// CreateSSH func for creates a new ssh connection.
// @Description Create a new ssh connection.
// @Summary create a new ssh connection
// @Tags SSH
// @Accept json
// @Produce json
// @Param Name body string true "Name"
// @Param Host body string true "Host"
// @Param Password body string true "Password"
// @Success 200 {object} models.SSH
// @Security ApiKeyAuth
// @Router /v1/ssh/create [post]
func CreateSSH(c *fiber.Ctx) error {
	return services.CreateSSHConnection(c)
}

func GetAllSSHConnections(c *fiber.Ctx) error {
	return services.GetAllSSHConnectionsByUserID(c)
}
