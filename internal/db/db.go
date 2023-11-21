package db

import (
	"context"
	"fmt"

	"cheezaram.tech/kataposis/internal/config"
	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func GetDB() *pgx.Conn {
	return db
}

func Connect(conf *config.DBConfig) (*pgx.Conn, error) {
	var err error

	dbURL, err := dbURL(conf)
	if err != nil {
		return nil, err
	}

	db, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dbURL(conf *config.DBConfig) (string, error) {
	dbConfig, err := config.NewConfig()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		dbConfig.PGUser,
		dbConfig.PGPass,
		dbConfig.PGAddr,
		dbConfig.PGDBName,
	), nil
}
