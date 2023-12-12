package constants

import "fmt"

// ======================= GENERAL ERRORS =======================
type ConnectionFailedError struct {
	Msg string
}

func (e ConnectionFailedError) Error() string { return e.Msg }

type InvalidConfigurationError struct {
	Msg string
}

func (e InvalidConfigurationError) Error() string { return e.Msg }

type EaseOffRetryError struct {
	Msg string
}

func (e EaseOffRetryError) Error() string { return e.Msg }

type NotInitializedError struct {
	Msg string
}

func (e NotInitializedError) Error() string { return e.Msg }

// ======================= FACTORIES ======================= 
var (
	ErrNotInitializedFactory = func(variable string) error {
		return NotInitializedError{
			Msg: fmt.Sprintf("trying to perform actions on uninitialized variable %v", variable),
		}
	}
	ErrInvalidConfigurationFactory = func(component string) error {
		return InvalidConfigurationError{
			Msg: fmt.Sprintf("invalid configuration for %s provided", component),
		}
	}
)