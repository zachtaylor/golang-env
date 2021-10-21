package env

import (
	"io/ioutil"
	"os"
	"strings"
)

// Values is a basic k/v map
type Values map[string]string

// New creates Values
func New() Values { return Values{} }

// Keys returns a new []string containing this Values' keys
func (s Values) Keys() []string {
	i, keys := 0, make([]string, len(s))
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

// Match returns a new Values, containing a subset of k/v from this Values
//
// When a key starts with prefix, the remainder of the key becomes the new key, for the unchanged value, in the new Values
func (s Values) Match(prefix string) Values {
	match, lpre := Values{}, len(prefix)
	for k, v := range s {
		if len(k) > lpre && k[:lpre] == prefix {
			match[k[lpre:]] = v
		}
	}
	return match
}

// Merge adds all k/v from another Values, prepend prefix for all new keys
func (s Values) Merge(prefix string, sub Values) Values {
	for k, v := range sub {
		s[prefix+k] = v
	}
	return s
}

// Parse parses "x=y" format as k/v to add to this Values
//
// All leading dash ('-') are removed, and "=y" defaults to "=true"
func (s Values) Parse(setting string) Values {
	for setting[0] == '-' {
		setting = setting[1:]
	}
	if kv := strings.Split(setting, "="); len(kv) == 1 {
		s[kv[0]] = "true"
	} else if len(kv) == 2 {
		s[kv[0]] = strings.Trim(kv[1], " \t")
	}
	return s
}

// ParseDefault is a macro for ParseDefaultFile, ParseEnv, and ParseFlags(os.Args[1:])
func (s Values) ParseDefault() (Values, error) {
	_, err := s.ParseDefaultFile()
	return s.ParseEnv().ParseArgs(os.Args[1:]), err
}

// MustParseDefault is a macro for MustParseDefaultFile, ParseEnv, and ParseFlags(os.Args[1:])
func (s Values) MustParseDefault() Values {
	return s.MustParseDefaultFile().ParseEnv().ParseArgs(os.Args[1:])
}

// ParseEnv scans os.Getenv for available updates to this Values
func (s Values) ParseEnv() Values {
	for _, k := range s.Keys() {
		if v := os.Getenv(k); len(v) > 1 {
			s[k] = v
		}
	}
	return s
}

// ParseFile parses a file line-by-line before calling `Parse`
//
// A pound sign ('#') comments the rest of the line
func (s Values) ParseFile(path string) (Values, error) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	for _, line := range strings.Split(string(file), "\n") {
		line = strings.Trim(strings.Split(line, "#")[0], " \r")
		if line != "" {
			s.Parse(line)
		}
	}
	return s, nil
}

// ParseDefaultFile returns ParseFile(DefaultFile)
func (s Values) ParseDefaultFile() (Values, error) { return s.ParseFile(DefaultFile) }

// MustParseFile returns ParseFile, with a panic for any error value
func (s Values) MustParseFile(path string) Values {
	if _, err := s.ParseFile(path); err != nil {
		panic(err.Error())
	}
	return s
}

// MustParseDefaultFile returns MustParseFile(DefaultFile)
func (s Values) MustParseDefaultFile() Values { return s.MustParseFile(DefaultFile) }

// ParseArgs adds []string encoded (e.g. os.Args[1:]) values to this Values
func (s Values) ParseArgs(args []string) Values {
	for _, arg := range args {
		s.Parse(arg)
	}
	return s
}
