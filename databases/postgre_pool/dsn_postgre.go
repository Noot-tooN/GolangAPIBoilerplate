package postgre_pool

import (
	"fmt"
	"strings"
)

type PgDSNBuilder struct {
	host        string
	port        string
	user        string
	password    string
	dbname      string
	sslmode     string // verify-ca, verify-full, disable
	sslcert     string // path
	sslkey      string // path
	sslrootcert string // path
}

func (dsnBuilder PgDSNBuilder) GenerateDsn() string {
	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v sslrootcert=%v sslkey=%v sslcert=%v",
		dsnBuilder.host,
		dsnBuilder.port,
		dsnBuilder.user,
		dsnBuilder.password,
		dsnBuilder.dbname,
		dsnBuilder.sslmode,
		dsnBuilder.sslrootcert,
		dsnBuilder.sslkey,
		dsnBuilder.sslcert,
	)
	
	return strings.TrimSpace(dns)
}

func NewPgDsnBuilder(host, port, user, password, dbName string, options ...PgDSNBuilderOption) (*PgDSNBuilder, error) {
	dsnBuilder := &PgDSNBuilder{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbName,
	}

	// Set default ssl mode to disabled
	// If there is an option for ssl mode the default ssl mode will be overriden
	WithSslModeFactory("disable")(dsnBuilder)

	for _, opt := range options {
		err := opt(dsnBuilder)
		if err != nil {
			return nil, err
		}
	}

	return dsnBuilder, nil
}
