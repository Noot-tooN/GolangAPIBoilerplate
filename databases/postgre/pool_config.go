package postgre

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgrePoolConfig struct {
	maxConns int32
	minConns int32
	maxConnLifetime time.Duration
	maxConnIdleTime time.Duration
	healthCheckPeriod time.Duration
	connectTimeout time.Duration
}

func NewPgPoolConfig(dsn PgDSNBuilder, options ...PostgrePoolConfigOption) (*pgxpool.Config, error) {
	pgClient := &PostgrePoolConfig{}

	// Set default values
	WithMaxConnsFactory(4)(pgClient)
	WithMinConnsFactory(0)(pgClient)
	WithMaxConnLifetimeFactory(time.Hour)(pgClient)
	WithMaxConnIdleTimeFactory(time.Minute * 30)(pgClient)
	WithHealthCheckPeriodTimeFactory(time.Minute)(pgClient)
	WithConnectTimeoutFactory(time.Second * 5)(pgClient)

	// Any user options will override the default values
	for _, opt := range options {
		err := opt(pgClient)
		if err != nil {
			return nil, err
		}
	}

	dbConfig, err := pgxpool.ParseConfig(dsn.GenerateDsn())

	if err!=nil {
		return nil, err
	}   

	return dbConfig, nil
}