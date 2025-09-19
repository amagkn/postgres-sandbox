package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) TripAll(ctx context.Context, input dto.TripAllInput) ([]dto.TripAllOutput, error) {
	trips, err := u.postgres.SelectManyTrips(ctx, input)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectManyTrips", err)
	}

	output := make([]dto.TripAllOutput, len(trips))
	for i := range trips {
		output[i] = dto.TripAllOutput(trips[i])
	}

	return output, nil
}
