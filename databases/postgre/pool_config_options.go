package postgre

import "time"

type PostgrePoolConfigOption func(*PostgrePoolConfig) error

func WithMaxConnsFactory(maxConns int32) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.maxConns = maxConns

		return nil
	}
}

func WithMinConnsFactory(minConns int32) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.minConns = minConns

		return nil
	}
}

func WithMaxConnLifetimeFactory(lifetime time.Duration) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.maxConnLifetime = lifetime

		return nil
	}
}

func WithMaxConnIdleTimeFactory(lifetime time.Duration) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.maxConnIdleTime = lifetime

		return nil
	}
}

func WithHealthCheckPeriodTimeFactory(period time.Duration) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.healthCheckPeriod = period

		return nil
	}
}

func WithConnectTimeoutFactory(timeout time.Duration) PostgrePoolConfigOption {
	return func(pcp *PostgrePoolConfig) error {
		pcp.connectTimeout = timeout

		return nil
	}
}