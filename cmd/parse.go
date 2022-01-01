package cmd

import (
	"os"

	"taylz.io/env"
)

// Parse1 parses os.Args for a single command name, remaining terms are parsed to overwrite given defaults
func Parse1(defaults env.Values) (string, env.Values) {
	if len(os.Args) < 2 {
		return "", defaults
	}
	if defaults == nil {
		defaults = env.Values{}
	}
	if len(os.Args) > 2 {
		env.SetWithArgs(defaults, os.Args[2:])
	}
	return os.Args[1], defaults
}
