package main

import (
	"log"

	"github.com/lmnq/jwt-token/config"
	"github.com/lmnq/jwt-token/internal/app"
)

func main() {
	// создание конфига
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	app.Run(cfg)
}
