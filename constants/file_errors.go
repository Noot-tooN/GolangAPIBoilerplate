package constants

import "fmt"

// ======================= FILE ERRORS =======================
type NoFileFoundError struct {
	Msg string
}

func (e NoFileFoundError) Error() string { return e.Msg }

// ======================= FACTORIES ======================= 
var (
	ErrNoFileFoundFactory = func(filePath string) error {
		return NoFileFoundError{
			Msg: fmt.Sprintf("could not find file at %v", filePath),
		}
	}
)