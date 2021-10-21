package env

// Version is the version number
const Version = "v0.0.2"

// DefaultFile is the file used by default
const DefaultFile = ".env"

// ParseDefault returns New().ParseDefault()
func ParseDefault() (Values, error) { return New().ParseDefault() }

// MustParseDefault returns New().MustParseDefault()
func MustParseDefault() Values { return New().MustParseDefault() }

// ParseDefaultFile returns New().ParseFile(DefaultFile)
func ParseDefaultFile() (Values, error) { return New().ParseDefaultFile() }

// ParseFile returns New().ParseFile(path)
func ParseFile(path string) (Values, error) { return New().ParseFile(path) }

// MustParseFile returns New().MustParseFile(path)
func MustParseFile(path string) Values { return New().MustParseFile(path) }

// ParseArgs returns New().ParseArgs(args)
func ParseArgs(args []string) Values { return New().ParseArgs(args) }
