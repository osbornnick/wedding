// Package db provides a helper for opening and verifying a PostgreSQL connection pool.
package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect creates a pgxpool connection pool using the supplied DSN and pings
// the server to confirm reachability.  A pool is used instead of a single
// connection so the HTTP server can safely handle concurrent requests.
func Connect(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	// Ping verifies that the DSN is valid and the database is reachable.
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}

	return pool, nil
}
