package gorm

import (
	"database/sql"
	"golangapi/constants"
	postgresqlclient "golangapi/databases/postgre_sql_client"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	defaultGormClient *gorm.DB
)

func NewPostgresGorm(pgdb *sql.DB, options ...gorm.Option) (*gorm.DB, error) {
	if len(options) == 0 {
		options = []gorm.Option{&gorm.Config{}}
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       pgdb,
		DriverName: constants.POSTGRES_DRIVER_NAME,
	}), options...)

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

func InitDefaultPostgresGorm(options ...gorm.Option) (*gorm.DB, error) {
	if len(options) == 0 {
		options = []gorm.Option{&gorm.Config{}}
	}

	gormDB, err := gorm.Open(postgres.New(
		postgres.Config{
			Conn: postgresqlclient.GetDefaultPostgreClient(),
		}),
		options...
	)

	if err != nil {
		return nil, err
	}

	defaultGormClient = gormDB

	return gormDB, nil
}

func GetDefaultGormClient() *gorm.DB {
	return defaultGormClient
}
