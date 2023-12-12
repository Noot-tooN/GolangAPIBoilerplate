package constants

// ======================= ENV ERRORS =======================
type EnvMissingValue struct {
	Msg string
}

func (e EnvMissingValue) Error() string { return e.Msg }

type EnvInvalidValue struct {
	Msg string
}

func (e EnvInvalidValue) Error() string { return e.Msg }

