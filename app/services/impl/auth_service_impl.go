package impl

import (
	"context"
	"errors"
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repositories"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/cache"
	"github.com/google/uuid"
)

type AuthService struct {
	repos repositories.Repositories
}

func NewAuthService(repos repositories.Repositories) services.Auth {
	return &AuthService{repos: repos}
}

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong email or password")
)

func (s *AuthService) SignUp(signUp *models.SignUp) (*models.User, error) {
	validate := utils.NewValidator()

	if err := validate.Struct(signUp); err != nil {
		return nil, err
	}

	role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		Email:        signUp.Email,
		PasswordHash: utils.GeneratePassword(signUp.Password),
		UserStatus:   1,
		UserRole:     role,
	}

	if err := validate.Struct(user); err != nil {
		return nil, err
	}

	if err := s.repos.User.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) SignIn(ctx context.Context, in *models.SignIn) (*services.Tokens, error) {
	user, err := s.repos.User.GetByEmail(in.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if ok := utils.ComparePasswords(user.PasswordHash, in.Password); !ok {
		return nil, ErrWrongPassword
	}

	creds, err := utils.GetCredentialsByRole(user.UserRole)
	if err != nil {
		return nil, err
	}

	t, err := utils.GenerateNewTokens(user.ID.String(), creds)
	if err != nil {
		return nil, err
	}

	connRedis, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}

	if err := connRedis.Set(ctx, user.ID.String(), t.Refresh, 0).Err(); err != nil {
		return nil, err
	}

	return &services.Tokens{
		Access:  t.Access,
		Refresh: t.Refresh,
	}, nil
}
