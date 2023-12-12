package postgresqlclient

import (
	"database/sql"
	"golangapi/config"
	"golangapi/constants"
	"golangapi/databases/postgre_dns"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/proullon/ramsql/driver"
)

var (
	defaultPGClient *sql.DB
)

type PostgreSqlClient struct {
	driver string
	dsn    string
	db     *sql.DB
}

func (psql *PostgreSqlClient) Connect() error {
	// Must use the pgx driver for postgre
	db, err := sql.Open(psql.driver, psql.dsn)

	if err != nil {
		psql.db = nil

		return constants.ConnectionFailedError{
			Msg: err.Error(),
		}
	}

	psql.db = db

	return nil
}

func (psql PostgreSqlClient) GetClient() (*sql.DB, error) {
	if psql.db == nil {
		return nil, constants.ErrNotInitializedFactory("psql db")
	}

	return psql.db, nil
}

func NewPostgreClient(driver string, dsn string) (*PostgreSqlClient, error) {
	pgClient := &PostgreSqlClient{
		driver: driver,
		dsn:    dsn,
	}

	return pgClient, nil
}

func SetDefaultPostgreClient(sqlCl PostgreSqlClient) error {
	if sqlCl.db == nil {
		return constants.ErrNotInitializedFactory("postgresql client")
	}

	defaultPGClient = sqlCl.db

	return nil
}

func SetDefaultPostgreClientFromDb(db *sql.DB) error {
	if db == nil {
		return constants.ErrNotInitializedFactory("postgresql client")
	}

	defaultPGClient = db

	return nil
}

func InitDefaultPostgreSqlClient(dnsOptions ...postgre_dns.PgDSNBuilderOption) error {
	dsnBuilder, err := postgre_dns.NewPgDsnBuilder(
		config.Config.Postgres.Host,
		strconv.Itoa(config.Config.Postgres.Port),
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.DbName,
	)

	if err != nil {
		return err
	}

	client, err := NewPostgreClient(
		constants.POSTGRES_DRIVER_NAME,
		dsnBuilder.GenerateDsn(),
	)

	if err != nil {
		return err
	}

	err = client.Connect()

	if err != nil {
		return err
	}

	defaultPGClient = client.db

	return nil
}

func GetDefaultPostgreClient() *sql.DB {
	return defaultPGClient
}
