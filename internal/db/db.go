package db

import (
	"context"
	"fmt"

	"cheezaram.tech/kataposis/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

// Connect creates a connection to the database using the configuration
// options specified by cfg.
func Connect(cfg *config.DBConfig) (*pgx.Conn, error) {
	url := dbURL(cfg)
	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.Connect failed")
	}
	return db, nil
}

func dbURL(cfg *config.DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.PGUser,
		cfg.PGPass,
		cfg.PGAddr,
		cfg.PGDBName,
	)
}
