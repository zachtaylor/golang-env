package env

import (
	"os"
	"strings"
)

// Values is a basic k/v map
type Values = map[string]string

// Keys returns a new []string containing the Values' keys
func Keys(env Values) []string {
	i, keys := 0, make([]string, len(env))
	for k := range env {
		keys[i] = k
		i++
	}
	return keys
}

// NewMatch returns a new Values, containing a subset of k/v from the Values
//
// A k/v pair is selected when the key has the prefix,
func NewMatch(env Values, prefix string) Values {
	match, lpre := Values{}, len(prefix)
	for k, v := range env {
		if len(k) > lpre && k[:lpre] == prefix {
			match[k[lpre:]] = v
		}
	}
	return match
}

// SetAll adds all k/v from sub to env, prepend prefix to all new keys
func SetAll(env, sub Values, prefix string) {
	for k, v := range sub {
		env[prefix+k] = v
	}
}

// SetWithLine parses "k=v" format to add to the Values
//
// lines without equals ('=') are implicitly "=true", i.e. Values[line]="true"
//
// all leading hyphen ('-') are removed
func SetWithLine(env Values, line string) {
	for line[0] == '-' {
		line = line[1:]
	}
	if before, after, ok := strings.Cut(line, "="); ok {
		env[before] = after
	} else {
		env[line] = "true"
	}
}

// SetDefault performs SetWithFile(DefaultFile), SetWithShell, and SetWithArgs(os.Args[1:])
//
// returns err from SetWithFile, doesn't block SetWithShell or SetWithArgs
func SetDefault(env Values) error {
	err := SetWithFile(env, DefaultFile)
	SetWithShell(env)
	SetWithArgs(env, os.Args[1:])
	return err
}

// SetWithFile reads a file, calling SetWithLine for all non-empty non-comment lines in the file
//
// A pound sign ('#') comments the rest of the line
func SetWithFile(env Values, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(file), "\n") {
		if line = strings.Trim(strings.Split(line, "#")[0], " \r"); line != "" {
			SetWithLine(env, line)
		}
	}
	return nil
}

// SetWithShell applies settings from os.Getenv to a Values
func SetWithShell(env Values) {
	for _, k := range Keys(env) {
		if v := os.Getenv(k); len(v) > 1 {
			env[k] = v
		}
	}
}

// SetWithArgs iterates args to call ParseLineWith
func SetWithArgs(env Values, args []string) {
	for _, arg := range args {
		SetWithLine(env, arg)
	}
}
