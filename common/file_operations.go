package common

import (
	"errors"
	"golangapi/constants"
	"os"
)

// If file exists no error is returned
// If file does not exist, error is returned
func FileExists(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return constants.ErrNoFileFoundFactory(path)
	}
	return nil
}
