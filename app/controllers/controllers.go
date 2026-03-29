package controllers

import (
	"github.com/create-go-app/fiber-go-template/app/services"
)

type Controllers struct {
	Auth  *AuthController
	Ssh   *SshController
	Token *TokenController
}

func NewControllers(services services.Services) Controllers {
	authCRL := NewAuthController(services)
	sshCRL := NewSshController(services)
	tokenCRL := NewTokenController(services)

	return Controllers{
		Auth:  authCRL,
		Ssh:   sshCRL,
		Token: tokenCRL,
	}
}
