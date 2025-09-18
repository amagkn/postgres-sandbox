package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) GetPassengersNames(ctx context.Context) (dto.GetPassengerNamesOutput, error) {
	passengers, err := u.postgres.SelectAllPassengers(ctx)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectAllPassengers", err)
	}

	output := make(dto.GetPassengerNamesOutput, len(passengers))
	for i, p := range passengers {
		output[i] = p.Name
	}

	return output, nil
}
