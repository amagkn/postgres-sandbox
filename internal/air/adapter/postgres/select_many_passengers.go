package postgres

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/internal/air/entity"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectManyPassengers(ctx context.Context, input dto.PassengerAllInput) ([]entity.Passenger, error) {
	// sql: SELECT id, name FROM air.passenger
	ds := goqu.From("air.passenger").Select("id", "name")
	if input.Suffix != "" {
		// ... where name like '%input.Suffix'
		ds = ds.Where(goqu.L("name LIKE ?", "%"+input.Suffix))
	}

	sql, _, err := ds.ToSQL()
	if err != nil {
		return nil, base_errors.WithPath("ds.ToSQL", err)
	}

	rows, err := p.pool.Query(ctx, sql)
	if err != nil {
		return nil, base_errors.WithPath("pool.Query", err)
	}
	defer rows.Close()

	var output []entity.Passenger
	for rows.Next() {
		var passenger entity.Passenger
		err := rows.Scan(&passenger.ID, &passenger.Name)
		if err != nil {
			return nil, base_errors.WithPath("rows.Scan", err)
		}

		output = append(output, passenger)
	}

	if err = rows.Err(); err != nil {
		return output, base_errors.WithPath("rows.Err", err)
	}

	return output, nil
}
