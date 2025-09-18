package postgres

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/entity"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectAllPassengers(ctx context.Context) ([]entity.Passenger, error) {
	// SELECT name FROM air.passenger
	ds := goqu.From("air.passenger").Select("name")

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
		err := rows.Scan(&passenger.Name)
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
