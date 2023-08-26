package repo

import (
	"context"

	"github.com/lmnq/jwt-token/internal/entity"
)

type Repo interface {
	StoreRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error
	GetRefreshToken(ctx context.Context, guid string) (*entity.RefreshToken, error)
}
