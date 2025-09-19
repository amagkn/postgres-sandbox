package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) PassengerAll(ctx context.Context, input dto.PassengerAllInput) ([]dto.PassengerAllOutput, error) {
	passengers, err := u.postgres.SelectManyPassengers(ctx, input)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectManyTrips", err)
	}

	output := make([]dto.PassengerAllOutput, len(passengers))
	for i := range passengers {
		output[i] = dto.PassengerAllOutput(passengers[i])
	}

	return output, nil
}
