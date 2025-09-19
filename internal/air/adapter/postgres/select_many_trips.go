package postgres

import (
	"context"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/internal/air/entity"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectManyTrips(ctx context.Context, input dto.TripsAllInput) ([]entity.Trip, error) {
	// sql: SELECT id, company_id, plane, town_from, town_to, time_out, time_in FROM air.trip where town_from = input.Town
	ds := goqu.
		From("air.trip").
		Select("id", "company_id", "plane", "town_from", "town_to", "time_out", "time_in").
		Where(goqu.C("town_from").Eq(input.Town))

	sql, _, err := ds.ToSQL()
	if err != nil {
		return nil, base_errors.WithPath("ds.ToSQL", err)
	}

	rows, err := p.pool.Query(ctx, sql)
	if err != nil {
		return nil, base_errors.WithPath("pool.Query", err)
	}
	defer rows.Close()

	var output []entity.Trip
	for rows.Next() {
		var trip entity.Trip
		err := rows.Scan(&trip.ID, &trip.CompanyID, &trip.Plane, &trip.TownFrom, &trip.TownTo, &trip.TimeOut, &trip.TimeIn)
		if err != nil {
			return nil, base_errors.WithPath("rows.Scan", err)
		}

		output = append(output, trip)
	}

	if err = rows.Err(); err != nil {
		return output, base_errors.WithPath("rows.Err", err)
	}

	return output, nil
}
