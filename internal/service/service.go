package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lmnq/jwt-token/internal/entity"
	"github.com/lmnq/jwt-token/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type TokenService struct {
	r repo.Repo
}

func New(r repo.Repo) *TokenService {
	return &TokenService{
		r: r,
	}
}

func (s *TokenService) CreateTokens(ctx context.Context, guid string) (*entity.Tokens, error) {
	accessToken, err := newAccessToken(guid)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken := newRefreshToken()
	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	err = s.r.StoreRefreshToken(ctx, &entity.RefreshToken{
		GUID:      guid,
		Hash:      string(refreshTokenHash),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		return nil, err
	}

	return &entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *TokenService) RefreshAccessToken(ctx context.Context, tokens *entity.Tokens) (*entity.AccessToken, error) {
	oldAccessToken, err := jwt.Parse(tokens.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing access jwt token: %w", err)
	}

	claims, ok := oldAccessToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing access jwt token claims: %w", err)
	}

	guid := claims["guid"].(string)

	refreshToken, err := s.r.GetRefreshToken(ctx, guid)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(refreshToken.Hash), []byte(tokens.RefreshToken)); err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		return nil, fmt.Errorf("refresh token expired")
	}

	accessToken, err := newAccessToken(guid)
	if err != nil {
		return nil, fmt.Errorf("error generating new access token: %w", err)
	}

	return &entity.AccessToken{
		Token: accessToken,
	}, nil
}
