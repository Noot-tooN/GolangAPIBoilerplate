package postgre

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	connectionPool *pgxpool.Pool
)

func InitPostgrePool(dsn PgDSNBuilder, options ...PostgrePoolConfigOption) error {
	config, err := NewPgPoolConfig(dsn, options...)

	if err != nil {
		return err
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), config)
	
	if err != nil {
		return err
	}

	connectionPool = connPool

	return nil
}

func GetPostgrePool() *pgxpool.Pool {
	return connectionPool
}