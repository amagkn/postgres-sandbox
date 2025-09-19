package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/internal/air/entity"
)

type Postgres interface {
	SelectManyPassengers(ctx context.Context, input dto.PassengerAllInput) ([]entity.Passenger, error)
	SelectManyTrips(ctx context.Context, input dto.TripAllInput) ([]entity.Trip, error)
	SelectPassengerNames(ctx context.Context) ([]string, error)
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
