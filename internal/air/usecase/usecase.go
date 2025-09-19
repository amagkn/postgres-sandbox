package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/internal/air/entity"
)

type Postgres interface {
	SelectManyPassengers(ctx context.Context) ([]entity.Passenger, error)
	SelectManyTrips(ctx context.Context, input dto.TripsAllInput) ([]entity.Trip, error)
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
