package kataposis

import (
	"context"
	"time"

	"cheezaram.tech/kataposis/internal/config"
	"cheezaram.tech/kataposis/internal/db"
	"github.com/jackc/pgx/v5"
)

// configValues is a function type that takes a pointer to a `config.DBConfig`.
// This is used to implement a functional options pattern for configuring the
// database connection.
type configValues func(cfg *config.DBConfig)

var (
	// pgDB is a global variable that stores the database connection.
	pgDB *pgx.Conn
	// cfg is a global variable that stores the database configuration.
	cfg *config.DBConfig
)

// LogEntry provides a fluent API for logging.
// It logs the `message` with the `resourceID`, `level` and `timestamp`.
//
// Only when the `Timestamp` method is called will the log will be entered
// into the database.
type LogEntry struct {
	resourceID logResourceID
	message    logMessage
	level      logLevel
	timestamp  time.Time
}

type (
	logResourceID string
	logMessage    string
	logLevel      string
	logTimestamp  time.Time
)

// Msg is used to set the message of the log entry. It takes in a `logMessage`
// and returns a `LogEntry` object.
//
// The log entry is never saved to the database until the `Timestamp` method
// is called.
func (l *LogEntry) Msg(msg logMessage) *LogEntry {
	if l == nil {
		panic("cannot call Msg on nil LogEntry")
	}

	l.message = msg
	return l
}

// RID sets the `resourceID` for the log entry. It takes in a `logResourceID`
// and returns a LogEntry object.
//
// The log entry is never saved to the database until the `Timestamp` method
// is called.
func (l *LogEntry) RID(rid logResourceID) *LogEntry {
	if l == nil {
		panic("cannot call RID on nil LogEntry")
	}

	l.resourceID = rid
	return l
}

// Level is used to set the log level of the entry. It takes in a logLevel and
// returns a LogEntry object.
//
// The log entry is never saved to the database until the `Timestamp` method
// is called.
func (l *LogEntry) Level(level logLevel) *LogEntry {
	if l == nil {
		panic("cannot call Level on nil LogEntry")
	}

	l.level = level
	return l
}

// Timestamp is used to set the timestamp of the entry. It takes in a
// logTimestamp and saves the entry to the database specified in the
// configuration.
//
// It returns an error if the database connection cannot be established.
func (l *LogEntry) Timestamp(ts time.Time) error {
	if l == nil {
		panic("cannot call Timestamp on nil LogEntry")
	}

	l.timestamp = ts

	var err error
	if pgDB, err = db.Connect(cfg); err != nil {
		return err
	}
	defer pgDB.Close(context.Background())

	return addLogEntry(l)
}

func WithPGUser(u string) configValues {
	return func(cfg *config.DBConfig) {
		cfg.PGUser = u
	}
}

func WithPGPassword(p string) configValues {
	return func(cfg *config.DBConfig) {
		cfg.PGPass = p
	}
}

func WithPGDB(d string) configValues {
	return func(cfg *config.DBConfig) {
		cfg.PGDBName = d
	}
}

// WithPGAddr sets the PostgreSQL connection address. Address should be in the
// format `host:port`. E.g. `localhost:5432`.
func WithPGAddr(addr string) configValues {
	return func(cfg *config.DBConfig) {
		cfg.PGAddr = addr
	}
}

// Init uses functional options pattern to initialize the database connection
// and create the logs table if it does not already exist.
// It returns a `LogEntry` object that can be used to log messages which are
// then saved to the database.
func Init(opts ...configValues) (*LogEntry, error) {
	cfg = &config.DBConfig{}

	for _, opt := range opts {
		opt(cfg)
	}

	var err error
	if pgDB, err = db.Connect(cfg); err != nil {
		return nil, err
	}
	defer pgDB.Close(context.Background())

	if _, err = pgDB.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS logs (
			id SERIAL PRIMARY KEY,
			rid TEXT,
			message TEXT,
			level TEXT,
			timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,
	); err != nil {
		return nil, err
	}

	return new(LogEntry), nil
}

func addLogEntry(l *LogEntry) error {
	_, err := pgDB.Exec(
		context.Background(),
		`INSERT INTO logs (
			rid, message, level, timestamp
		) VALUES ($1, $2, $3, $4)`,
		l.resourceID, l.message, l.level, l.timestamp,
	)

	return err
}
