package postgres

import (
	"context"

	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectPassengerNames(ctx context.Context) ([]string, error) {
	// sql: SELECT name FROM air.passenger
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

	var output []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, base_errors.WithPath("rows.Scan", err)
		}
		output = append(output, name)
	}

	if err = rows.Err(); err != nil {
		return nil, base_errors.WithPath("rows.Err", err)
	}

	return output, nil
}
