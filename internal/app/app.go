package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/amagkn/postgres-sandbox/config"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/amagkn/postgres-sandbox/pkg/http_server"
	"github.com/amagkn/postgres-sandbox/pkg/logger"
	"github.com/amagkn/postgres-sandbox/pkg/postgres"
	"github.com/amagkn/postgres-sandbox/pkg/router"
	"github.com/go-chi/chi/v5"
)

type Dependences struct {
	RouterHTTP *chi.Mux
	Postgres   *postgres.Pool
}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependences

	deps.Postgres, err = postgres.New(ctx, c.Postgres)
	if err != nil {
		return base_errors.WithPath("postgres.New", err)
	}
	defer deps.Postgres.Close()

	deps.RouterHTTP = router.New()

	AirDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	logger.Info("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		logger.Info("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		logger.Error(err, "App got notify: httpServer.Notify")
	}

	logger.Info("App is stopping")
}
