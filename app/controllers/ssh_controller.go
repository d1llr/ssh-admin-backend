package controllers

import (
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/gofiber/fiber/v2"
)

// CreateSSH method to create new SSH connection
// @Description Create new ssh connection.
// @Summary creating new ssh connection
// @Tags SSH
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param host body string true "Host"
// @Param password body string true "Password"
// @Success 200 {object} models.SSH
// @Router /v1/user/sign/up [post]
func CreateSSH(c *fiber.Ctx) error {
	return services.CreateSSHConnection(c)
}
