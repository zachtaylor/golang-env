package env

import (
	"io/ioutil"
	"os"
	"strings"
)

// Values is a basic k/v map
type Values map[string]string

// New creates `Values`
func New() Values { return Values{} }

// Keys returns a new `[]string` containing this `Values` keys
func (s Values) Keys() []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

// Match returns a new `Values`, containing k/v from this `Values`, where `k` begins with `prefix`, and `k` in the new `Values` has `prefix` removed
func (s Values) Match(prefix string) Values {
	match, lpre := Values{}, len(prefix)
	for k, v := range s {
		if len(k) > lpre && k[:lpre] == prefix {
			match[k[lpre:]] = v
		}
	}
	return match
}

// Merge writes another `Values` data into this `Values`, adding `prefix` before each new key
func (s Values) Merge(prefix string, sub Values) Values {
	for k, v := range sub {
		s[prefix+k] = v
	}
	return s
}

// Parse parses `"x=y"` format to add a k/v to this `Values`
//
// `=y` is optional, defaults to `=true`
func (s Values) Parse(setting string) Values {
	if kv := strings.Split(setting, "="); len(kv) == 1 {
		s[kv[0]] = "true"
	} else if len(kv) == 2 {
		s[kv[0]] = strings.Trim(kv[1], " \t")
	}
	return s
}

// ParseDefault is a macro for ParseDefaultFile, ParseEnv, and ParseFlags(os.Args[1:])
func (s Values) ParseDefault() (Values, error) {
	s, err := s.ParseDefaultFile()
	return s.ParseEnv().ParseArgs(os.Args[1:]), err
}

// MustParseDefault calls ParseDefault, with a panic for any error value
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

// ParseDefaultFile calls ParseFile with ".env"
func (s Values) ParseDefaultFile() (Values, error) { return s.ParseFile(".env") }

// MustParseFile calls ParseFile, with a panic for any error value
func (s Values) MustParseFile(path string) Values {
	if _, err := s.ParseFile(path); err != nil {
		panic(err.Error())
	}
	return s
}

// MustParseDefaultFile calls MustParseFile with ".env"
func (s Values) MustParseDefaultFile() Values { return s.MustParseFile(".env") }

// ParseArgs formats each arg before calling `Parse`
//
// args e.g. `os.Args[1:]`
//
// an arg that begins with a hypen (`-`) hyphen is removed
//
// an arg that ends with a backslash (`\`) is treated as control escape for space concatenation
func (s Values) ParseArgs(args []string) Values {
	combargs := make([]string, 0)
	for i := 0; i < len(args); i++ {
		str := args[i]
		if len(str) < 1 {
			continue
		}
		for str[0] == '-' {
			str = str[1:]
		}
		for str[len(str)-1] == '\\' && i+1 < len(args) {
			i++
			str = str[:len(str)-1] + " " + args[i]
		}
		combargs = append(combargs, str)
	}
	for _, arg := range combargs {
		s.Parse(arg)
	}
	return s
}
