package app

import (
	"github.com/amagkn/postgres-sandbox/internal/air/adapter/postgres"
	"github.com/amagkn/postgres-sandbox/internal/air/controller/http_router"
	"github.com/amagkn/postgres-sandbox/internal/air/usecase"
)

func AirDomain(d Dependences) {
	airUseCase := usecase.New(postgres.New(d.Postgres))

	http_router.AirRoutes(d.RouterHTTP, airUseCase)
}
