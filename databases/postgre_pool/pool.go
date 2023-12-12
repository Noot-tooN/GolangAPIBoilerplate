package postgre_pool

import (
	"context"
	"golangapi/config"
	"log"
	"strconv"

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

func InitDefaultPool() error {
	dsnBuilder, err := NewPgDsnBuilder(
		config.Config.Postgres.Host,
		strconv.Itoa(config.Config.Postgres.Port),
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.DbName,
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = InitPostgrePool(*dsnBuilder)

	if err != nil {
		log.Fatalln(err)
	}

	pool := GetPostgrePool()

	conn, err := pool.Acquire(context.Background())

	defer conn.Release()

	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func GetPostgrePool() *pgxpool.Pool {
	return connectionPool
}