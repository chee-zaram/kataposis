package db

import (
	"context"
	"fmt"

	"cheezaram.tech/kataposis/internal/config"
	"github.com/jackc/pgx/v5"
)

// Connect creates a connection to the database using the configuration
// options specified by cfg.
func Connect(cfg *config.DBConfig) (*pgx.Conn, error) {
	url := dbURL(cfg)
	return pgx.Connect(context.Background(), url)
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
