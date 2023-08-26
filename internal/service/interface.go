package service

import (
	"context"

	"github.com/lmnq/jwt-token/internal/entity"
)

type Service interface {
	CreateTokens(ctx context.Context, guid string) (*entity.Tokens, error)
	RefreshAccessToken(ctx context.Context, tokens *entity.Tokens) (*entity.AccessToken, error)
}
