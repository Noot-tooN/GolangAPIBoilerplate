package postgre_dns

import (
	"golangapi/common"
	"golangapi/constants"
)

var (
	sslModes = map[constants.PostgreSSLMode]string{
		constants.SSLModeVerifyCa:   constants.SSLModeVerifyCa,
		constants.SSLModeVerfiyFull: constants.SSLModeVerfiyFull,
		constants.SSLModeDisable:    constants.SSLModeDisable,
	}
)

type PgDSNBuilderOption func(*PgDSNBuilder) error

func WithSslModeFactory(mode constants.PostgreSSLMode) PgDSNBuilderOption {
	val, ok := sslModes[mode]

	// Default to disable
	if !ok {
		return func(psc *PgDSNBuilder) error {
			psc.sslmode = "disable"

			return nil
		}
	}
	
	return func(psc *PgDSNBuilder) error {
		psc.sslmode = val

		return nil
	}
}

func WithSslCertFactory(path string) PgDSNBuilderOption {
	return func(psc *PgDSNBuilder) error {
		err := common.FileExists(path)

		if err != nil {
			return constants.ErrInvalidConfigurationFactory("postgre client ssl cert")
		}

		psc.sslcert = path

		return nil
	}
}

func WithSslKeyFactory(path string) PgDSNBuilderOption {
	return func(psc *PgDSNBuilder) error {
		err := common.FileExists(path)

		if err != nil {
			return constants.ErrInvalidConfigurationFactory("postgre client ssl key")
		}

		psc.sslkey = path

		return nil
	}
}

func WithSslRootCertFactory(path string) PgDSNBuilderOption {
	return func(psc *PgDSNBuilder) error {
		err := common.FileExists(path)

		if err != nil {
			return constants.ErrInvalidConfigurationFactory("postgre client ssl root cert")
		}

		psc.sslrootcert = path

		return nil
	}
}
