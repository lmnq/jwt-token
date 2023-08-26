package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lmnq/jwt-token/internal/entity"
	"github.com/lmnq/jwt-token/internal/service"
	"golang.org/x/exp/slog"
)

type tokenController struct {
	sl *slog.Logger
	s  service.Service
}

func newTokenController(sl *slog.Logger, s service.Service) *tokenController {
	return &tokenController{sl, s}
}

func (c *tokenController) createTokens(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	tokens, err := c.s.CreateTokens(r.Context(), guid)
	if err != nil {
		c.sl.Error("error creating tokens: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tokens)
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (c *tokenController) refreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var tokens entity.Tokens

	err := json.NewDecoder(r.Body).Decode(&tokens)
	if err != nil {
		c.sl.Error("error decoding tokens: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessToken, err := c.s.RefreshAccessToken(r.Context(), &tokens)
	if err != nil {
		c.sl.Error("error refreshing access token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessTokenResponse{
		AccessToken: accessToken.Token,
	})
}
