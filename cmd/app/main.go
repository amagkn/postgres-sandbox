package main

import (
	"context"

	"github.com/amagkn/postgres-sandbox/config"
	"github.com/amagkn/postgres-sandbox/internal/app"
	"github.com/amagkn/postgres-sandbox/pkg/logger"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(c.Logger)

	err = app.Run(ctx, c)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
