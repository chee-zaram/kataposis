package config

import (
	"errors"
	"os"
)

type DBConfig struct {
	PGUser   string
	PGPass   string
	PGAddr   string
	PGDBName string
}

const (
	pgUser   = "KTPSS_PG_USER"
	pgPass   = "KTPSS_PG_PASS"
	pgAddr   = "KTPSS_PG_ADDR"
	pgDBName = "KTPSS_PG_DB"
)

// NewConfig returns a configuration struct based on the current values of the
// environment variables.
func NewConfig() (*DBConfig, error) {
	user, ok := os.LookupEnv(pgUser)
	if !ok {
		return nil, errors.New("database user not set")
	}

	pass, ok := os.LookupEnv(pgPass)
	if !ok {
		return nil, errors.New("database password not set")
	}

	addr, ok := os.LookupEnv(pgAddr)
	if !ok {
		return nil, errors.New("database address not set")
	}

	db, ok := os.LookupEnv(pgDBName)
	if !ok {
		return nil, errors.New("database name not set")
	}

	return &DBConfig{
		PGUser:   user,
		PGPass:   pass,
		PGAddr:   addr,
		PGDBName: db,
	}, nil
}
