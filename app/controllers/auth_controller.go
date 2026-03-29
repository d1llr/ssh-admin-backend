package controllers

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/cache"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	services services.Services
}

func NewAuthController(services services.Services) *AuthController {
	return &AuthController{services: services}
}

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func (h *AuthController) UserSignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := c.BodyParser(signUp); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	user, err := h.services.Auth.SignUp(signUp)
	if err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	return utils.NewResponse(
		c,
		fiber.StatusCreated,
		false,
		"User created successfully",
		fiber.Map{
			"user": user,
		},
	)
}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Success 200 {string} status "ok"
// @Router /v1/user/sign/in [post]
func (h *AuthController) UserSignIn(c *fiber.Ctx) error {
	signIn := &models.SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	tokens, err := h.services.Auth.SignIn(c.Context(), signIn)
	if err != nil {
		return utils.NewResponse(
			c,
			fiber.StatusBadRequest,
			true,
			err.Error(),
		)
	}

	return utils.NewResponse(
		c,
		fiber.StatusOK,
		false,
		"singin success",
		fiber.Map{
			"tokens": tokens,
		},
	)
}

// UserSignOut method to de-authorize user and delete refresh token from Redis.
// @Description De-authorize user and delete refresh token from Redis.
// @Summary de-authorize user and delete refresh token from Redis
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user/sign/out [post]
func UserSignOut(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Define user ID.
	userID := claims.UserID.String()

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Save refresh token to Redis.
	errDelFromRedis := connRedis.Del(context.Background(), userID).Err()
	if errDelFromRedis != nil {
		// Return status 500 and Redis deletion error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errDelFromRedis.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
