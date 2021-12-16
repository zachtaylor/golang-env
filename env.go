package env

// Version is the version number
const Version = "v0.1.0"

// DefaultFile is the file used by default
const DefaultFile = ".env"

// ParseDefault returns the Values from default sources
func ParseDefault() (Values, error) { return New().ParseDefault() }

// MustParseDefault calls ParseDefault and panics on error
func MustParseDefault() Values { return New().MustParseDefault() }

// ShouldParseDefault calls ParseDefault and drops any file error
func ShouldParseDefault() Values { return New().ShouldParseDefault() }

// ParseDefaultFile returns the Values for the default file
func ParseDefaultFile() (Values, error) { return New().ParseDefaultFile() }

// ParseFile returns the Values from the given file path
func ParseFile(path string) (Values, error) { return New().ParseFile(path) }

// MustParseFile calls ParseFile and panics on error
func MustParseFile(path string) Values { return New().MustParseFile(path) }

// ShouldParseFile calls ParseFile and drops any file error
func ShouldParseFile(path string) Values { return New().ShouldParseFile(path) }

// ParseArgs returns Values from the given arguments (e.g. os.Args[1:])
func ParseArgs(args []string) Values { return New().ParseArgs(args) }
