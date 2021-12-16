package env

// Builder is a Values that exposes parse access
type Builder map[string]string

// New creates a Builder
func New() Builder { return Builder{} }

// Match returns a new Values, containing a subset of k/v from this Values
//
// see also NewMatch
func (env Builder) Match(prefix string) Values { return NewMatch(env, prefix) }

// Merge adds all k/v from sub to this Values, prepend prefix on all new keys
//
// see also SetAll
func (env Builder) Merge(prefix string, sub Values) { SetAll(env, sub, prefix) }

// ParseDefault loads DefaultFile, sets values using os.Getenv, and SetWithArgs(os.Args[1:])
func (env Builder) ParseDefault() (Builder, error) { return env, SetDefault(env) }

// MustParseDefault calls ParseDefault and panics on error
func (env Builder) MustParseDefault() Builder { return must(env.ParseDefault()) }

// ShouldParseDefault calls ParseDefault and ignores file error
func (env Builder) ShouldParseDefault() Builder { return should(env.ParseDefault()) }

// ParseDefaultFile loads DefaultFile
func (env Builder) ParseDefaultFile() (Builder, error) { return env, SetWithFile(env, DefaultFile) }

// ParseFile loads a file and sets all values defined in .profile style
func (env Builder) ParseFile(path string) (Builder, error) { return env, SetWithFile(env, path) }

// MustParseFile calls ParseFile and panics on error
func (env Builder) MustParseFile(path string) Builder { return must(env.ParseFile(path)) }

// ShouldParseFile calls ParseFile and drops any file error
func (env Builder) ShouldParseFile(path string) Builder { return should(env.ParseFile(path)) }

// ParseShell looks up all keys in os.Getenv, sets new values
func (env Builder) ParseShell() Builder {
	SetWithShell(env)
	return env
}

// ParseArgs sets all lines within args
func (env Builder) ParseArgs(args []string) Builder {
	SetWithArgs(env, args)
	return env
}

// must handles error returns with a panic
func must(env Builder, err error) Builder {
	if err != nil {
		panic(err)
	}
	return env
}

// should ignores error returns
func should(env Builder, err error) Builder { return env }
