package env

// Version is the version number
const Version = "v0.0.0"

// ParseDefault returns a new Service, loaded with `ParseDefault()`
func ParseDefault() Values { return New().ParseDefault() }

// ParseDefaultFile returns a new Service, loaded with `ParseDefaultFile()`
func ParseDefaultFile() Values { return New().ParseDefaultFile() }

// ParseFile returns a new Service, loaded with `ParseFile(path)`
func ParseFile(path string) (Values, error) { return New().ParseFile(path) }

// ParseArgs returns a new Service, loaded with `ParseArgs()`
func ParseArgs(args []string) Values { return New().ParseArgs(args) }
