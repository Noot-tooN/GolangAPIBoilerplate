package constants

import "fmt"

// ======================= FILE ERRORS =======================

type NoFileFoundError struct {
	Msg string
}

func (e NoFileFoundError) Error() string { return e.Msg }

// ======================= GENERAL ERRORS ======================= 
type InvalidConfigurationError struct {
	Msg string
}

func (e InvalidConfigurationError) Error() string { return e.Msg }

type EaseOffRetryError struct {
	Msg string
}

func (e EaseOffRetryError) Error() string { return e.Msg }

var (
	ErrNoFileFoundFactory = func(filePath string) error {
		return NoFileFoundError{
			Msg: fmt.Sprintf("could not find file at %v", filePath),
		}
	}
	ErrInvalidConfigurationFactory = func(component string) error {
		return InvalidConfigurationError{
			Msg: fmt.Sprintf("invalid configuration for %s provided", component),
		}
	}

)