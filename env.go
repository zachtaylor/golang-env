package env

// Version is the version number
const Version = "v0.0.1"

// ParseDefault returns a new Values, loaded by ParseDefault
func ParseDefault() (Values, error) { return New().ParseDefault() }

// MustParseDefault returns a new Values, loaded by MustParseDefault
func MustParseDefault() Values { return New().MustParseDefault() }

// ParseDefaultFile returns a new Values, loaded with `ParseDefaultFile()`
func ParseDefaultFile() (Values, error) { return New().ParseDefaultFile() }

// ParseFile returns a new Values, loaded with `ParseFile(path)`
func ParseFile(path string) (Values, error) { return New().ParseFile(path) }

// ParseArgs returns a new Values, loaded with `ParseArgs()`
func ParseArgs(args []string) Values { return New().ParseArgs(args) }
