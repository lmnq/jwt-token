package service

import (
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	secretKey = "any-secret-key"
)

func newAccessToken(guid string) (string, error) {
	accessClaims := jwt.MapClaims{
		"guid":    guid,
		"expires": time.Now().Add(time.Minute * 30).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)

	accessString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessString, nil
}

func newRefreshToken() string {
	refreshToken := uuid.New().String()
	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	return refreshTokenBase64
}
