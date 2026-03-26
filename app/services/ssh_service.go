package services

import (
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateSSHConnection(c *fiber.Ctx) error {
	ssh := &models.SSH{}

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusInternalServerError,
			true,
			err.Error(),
		)
	}

	if expErr := utils.Check_exp(c, *claims); expErr != nil {
		return expErr
	}

	if err := c.BodyParser(ssh); err != nil {
		// Return status 400 and error message.
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusInternalServerError,
			true,
			err.Error(),
		)
	}

	validate := utils.NewValidator()

	ssh.ID = uuid.New()
	ssh.CreatedAt = time.Now()
	ssh.UserID = claims.UserID

	if err := validate.Struct(ssh); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}
	if err := db.CreateSSHConnection(ssh); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusInternalServerError,
			true,
			err.Error(),
		)
	}

	return utils.NewResponse(
		c,
		fiber.StatusCreated,
		false,
		"Connection created successfully",
		fiber.Map{
			"ssh": ssh,
		},
	)
}
