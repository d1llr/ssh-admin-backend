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

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return utils.NewResponse(
			c,
			fiber.StatusInternalServerError,
			true,
			err.Error(),
		)
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Set initialized default data for book:
	ssh.ID = uuid.New()
	ssh.CreatedAt = time.Now()
	ssh.UserID = claims.UserID

	// Validate book fields.
	if err := validate.Struct(ssh); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create book by given model.
	if err := db.create(ssh); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
}
