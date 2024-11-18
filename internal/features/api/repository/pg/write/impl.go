package write

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WritePostgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (*WritePostgres, error) {
	if db == nil {
		return nil, fmt.Errorf("pgxpool is nil")
	}
	return &WritePostgres{
		db: db,
	}, nil
}
