package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) TripsAll(ctx context.Context, input dto.TripsAllInput) ([]dto.TripsAllOutput, error) {
	trips, err := u.postgres.SelectManyTrips(ctx, input)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectManyTrips", err)
	}

	output := make([]dto.TripsAllOutput, len(trips))
	for i := range trips {
		output[i] = dto.TripsAllOutput(trips[i])
	}

	return output, nil
}
