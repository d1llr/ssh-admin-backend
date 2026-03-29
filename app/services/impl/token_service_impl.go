// app/services/impl/token_service_impl.go
package impl

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/repositories"
	svc "github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/cache"
	"github.com/google/uuid"
)

type TokenService struct {
	repos repositories.Repositories
}

func NewTokenService(r repositories.Repositories) svc.Token {
	return &TokenService{repos: r}
}

func (s *TokenService) RenewToken(ctx context.Context, userId uuid.UUID) (*svc.Tokens, error) {
	user, err := s.repos.User.GetByID(userId)
	if err != nil {
		return nil, utils.CreateError(utils.UserNotFoundError)
	}

	creds, err := utils.GetCredentialsByRole(user.UserRole)
	if err != nil {
		return nil, err
	}

	t, err := utils.GenerateNewTokens(userId.String(), creds)
	if err != nil {
		return nil, err
	}

	connRedis, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}

	if err := connRedis.Set(ctx, userId.String(), t.Refresh, 0).Err(); err != nil {
		return nil, err
	}

	return &svc.Tokens{
		Access:  t.Access,
		Refresh: t.Refresh,
	}, nil
}
