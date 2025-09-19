package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) PassengerNames(ctx context.Context) (dto.PassengerNamesOutput, error) {
	passengers, err := u.postgres.SelectManyPassengers(ctx)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectManyPassengers", err)
	}

	output := make(dto.PassengerNamesOutput, len(passengers))
	for i, p := range passengers {
		output[i] = p.Name
	}

	return output, nil
}
