package common

import (
	"fmt"
	"golangapi/constants"
	"time"
)

// Attempt to execute function 'f' at maximum 'attempts' number of times
// Sleep the thread for sleep duration
// Each consequent attempt will sleep double than the previous attempt
// If function 'f' wishes to break the retry attempts the function should return for the bool return value
func EaseOffRetryStrategy(attempts int, sleep time.Duration, f func() (bool, error)) (err error) {
	var shouldRetry bool

	for i := 0; i < attempts; i++ {
		if i > 0 {
			time.Sleep(sleep)
			sleep *= 2
		}

		shouldRetry, err = f()

		if !shouldRetry {
			return err
		}
	}
	
	if err != nil {
		return constants.EaseOffRetryError{
			Msg: fmt.Sprintf("after %d attempts, last error: %v", attempts, err),
		}
	}
	
	return nil
}
