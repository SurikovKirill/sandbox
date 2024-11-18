package read

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReadPostgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (*ReadPostgres, error) {
	if db == nil {
		return nil, fmt.Errorf("pgxpool is nil")
	}
	return &ReadPostgres{
		db: db,
	}, nil
}

func (ReadPostgres) Create(context.Context) error {
	return nil
}

func (ReadPostgres) Update(context.Context) error {
	return nil
}
func (ReadPostgres) Delete(context.Context) error {
	return nil
}
