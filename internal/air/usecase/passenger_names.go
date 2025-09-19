package usecase

import (
	"context"

	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
)

func (u *UseCase) PassengerNames(ctx context.Context) ([]string, error) {
	passengerNames, err := u.postgres.SelectPassengerNames(ctx)
	if err != nil {
		return nil, base_errors.WithPath("u.postgres.SelectPassengerNames", err)
	}

	return passengerNames, nil
}
