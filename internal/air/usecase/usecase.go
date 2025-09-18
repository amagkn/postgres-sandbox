package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/entity"
)

type Postgres interface {
	SelectAllPassengers(ctx context.Context) ([]entity.Passenger, error)
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
