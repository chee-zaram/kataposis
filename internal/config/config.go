package config

type DBConfig struct {
	PGUser   string
	PGPass   string
	PGAddr   string
	PGDBName string
}

// NewEmptyDBConfig returns a new empty database configuration struct.
func NewEmptyDBConfig() *DBConfig {
	return &DBConfig{}
}
