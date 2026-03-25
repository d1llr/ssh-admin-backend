package controllers

import (
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}
