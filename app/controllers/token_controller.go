// controllers/token_controller.go
package controllers

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/pkg/httpctx"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type TokenController struct {
	services services.Services
}

func NewTokenController(services services.Services) *TokenController {
	return &TokenController{services: services}
}

func (t *TokenController) RenewTokens(c *fiber.Ctx) error {
	userId := httpctx.UserID(c)

	tokens, err := t.services.Token.RenewToken(context.Background(), userId)
	if err != nil {
		return utils.NewResponse(c, fiber.StatusInternalServerError, true, err.Error())
	}

	return utils.NewResponse(
		c,
		fiber.StatusOK,
		false,
		"Succesfully generated new tokens pair",
		fiber.Map{
			"tokens": tokens,
		},
	)
}
