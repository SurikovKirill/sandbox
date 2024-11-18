package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnectionPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.TODO(), "")
	if err != nil {
		return nil, fmt.Errorf("failed to make new pgxpool connection: %w", err)
	}

	return pool, nil
}
