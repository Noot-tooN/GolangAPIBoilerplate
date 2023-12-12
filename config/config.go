package config

import (
	"golangapi/constants"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

// Configuration fields will be loaded from (in order):
//
//  1. defaults set in structure tags
//  2. loaded from files:
//     whichever file has the value first, that value will be set
//  3. from corresponding environment variables with the set prefix if any
//  4. command-line flags with the set prefix if any
type ServerConfig struct {
	Postgres struct {
		Host     string `default:"127.0.0.1" env:"DB_HOST" flag:"host" yaml:"postgres.host" usage:"Set an url to where the PostgreSQL database is hosted"`
		Port     int    `default:"5432" env:"DB_PORT" flag:"port" yaml:"postgres.port" usage:"Set a port on which the PostgreSQL database is listening"`
		User     string `default:"custom-user" env:"DB_USER" flag:"user" yaml:"postgres.user" usage:"Set a user for PostgreSQL database"`
		Password string `default:"custom-pass" env:"DB_PASSWORD" flag:"pass" yaml:"postgres.pass" usage:"Set a password which will be user for authentication for PostgreSQL database"`
		DbName   string `default:"custom-db-name" env:"DB_NAME" flag:"dbname" yaml:"postgres.dbname" usage:"Set a PostgreSQL database name that is going to be targeted"`
	}
	Server struct {
		Host     string `default:"localhost" env:"HOST" flag:"host" yaml:"server.host" usage:"Set a host that the server will listen on"`
		Port     int    `default:"9911" env:"PORT" flag:"port" yaml:"server.port" usage:"Set a port number that the server will listen on"`
	}
}

type GlobalConfig struct {
	Env string
	ServerConfig
}

func (gc GlobalConfig) IsProd() bool {
	return gc.Env == constants.ProductionEnv
}

var Config GlobalConfig

func InitConfig(osArgs []string) error {
	// Load in the env
	env, newArgs, err := ReadEnvConfig(osArgs)
	if err != nil {
		return err
	}
	// Load in the server config
	var serverConfig ServerConfig
	serverLoader := aconfig.LoaderFor(&serverConfig, aconfig.Config{
		AllowUnknownFields: false,
		AllowUnknownEnvs:   false,
		AllowUnknownFlags:  false,
		SkipFlags:          false,
		Args:               newArgs,
		Files: []string{
			".env",
			"config.yaml",
		},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
			".env":  aconfigdotenv.New(),
		},
	})
	if err := serverLoader.Load(); err != nil {
		return err
	}
	// Merge configs
	Config = GlobalConfig{
		Env:          env,
		ServerConfig: serverConfig,
	}
	return nil
}
