package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lmnq/jwt-token/internal/service"
	"golang.org/x/exp/slog"
)

func NewRouter(router *chi.Mux, sl *slog.Logger, s service.Service) {
	// middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	// controller
	tc := newTokenController(sl, s)

	// routes
	router.Route("/", func(r chi.Router) {
		r.Post("/tokens", tc.createTokens)
		r.Post("/refresh", tc.refreshAccessToken)
	})
}
