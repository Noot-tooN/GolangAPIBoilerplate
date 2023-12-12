package config

import (
	"fmt"
	"golangapi/constants"
)

// Searches the provided arguments and returns the found env string, position in argument array and no error
// If there was any kind of error it returns empty string, -1 and the error
func getEnvValue(args []string) (string, int, error) {
	// Default case
	if len(args) == 0 {
		return constants.DefaultEnv, -1, nil
	}

	var pos int = -1

	for i, arg := range args {
		if arg == "-env" || arg == "--env" {
			pos = i
			break
		}
	}

	// Didn't find -env or --env, default value
	if pos == -1 {
		return constants.DefaultEnv, -1, nil
	}

	// Out of bounds just --env provided
	if pos+1 >= len(args) {
		return "", -1, constants.EnvMissingValue{
			Msg: "env flag is present but value was not set",
		}
	}

	envVal := args[pos+1]
	val, ok := constants.AllowedEnvs[envVal]

	// Not allowed value
	if !ok {
		keys := make([]string, 0, len(constants.AllowedEnvs))

		for k := range constants.AllowedEnvs {
			keys = append(keys, k)
		}

		return "", -1, constants.EnvMissingValue{
			Msg: fmt.Sprintf("invalid env value provided: %v\nallowed values %v", envVal, keys),
		}
	}

	return val, pos, nil
}

func ReadEnvConfig(args []string) (string, []string, error) {
	val, pos, err := getEnvValue(args)
	if err != nil {
		return "", nil, err
	}
	if pos == -1 {
		return val, args, nil
	}
	return val, append(args[:pos], args[pos+2:]...), nil
}
