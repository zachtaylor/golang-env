# taylz.io/env

env provides a simple runtime environment, `map[string]string` as `env.Values`

Data is provided by flags, files, and/or environment variables

## Install

```sh
go get taylz.io/env
```

## Philosophy

The goal of this library is to provide a SIMPLE consistent local environment header for potentially complex systems in development

It is an explicit non-goal to support [client](https://direnv.net/) [workflow](https://github.com/hyperupcall/autoenv) via [system](https://stackoverflow.com/questions/45216663/how-to-automatically-activate-virtualenvs-when-cding-into-a-directory) [tooling](https://github.com/ashald/EnvFile)

## Comparison

### spf13/viper

Advantages of [spf13/viper](https://github.com/spf13/viper):

- Mature
- Multiple file types
- Remote k/v readers
- Realtime update of runtime settings
- Type conversion
- Supports [spf13/cobra](https://github.com/spf13/cobra)

Advantages of taylz.io/env:

- Simple
- No external dependencies

## Parsing

### Flags

Instead of using stdlib `flag`, this library uses `os.Args` to receive arbitrary k/v

This library takes the opinion that leading hyphen (`-`) within keys can be confusing for clients' casual users. Flag keys, therefore, are interpreted with all leading hyphen removed; flags (os.Args) are interpreted as in the form `"[-]*{K}[={V}]"` where `{K}, {V}` are string values, and `{V}` defaults to `true`.

### Files

The file format is assumed to be a basic shell-like format

```sh
# .env file
PORT=:443
SSH_KEY=/private/ssh/key
SSH_CERT=/private/ssh/cert
```

## Examples

This simple example has 1 environment setting. That is `"PORT"`. The default value is `":8080"`

`env.ParseDefault` will parse ".env", call `os.Getenv` with each key to check for 

```go
package main 

import "taylz.io/env"

func main() {
	// parse ".env" file, os env, and cli args
	env, _ := defaultEnv().ParseDefault()
	// access runtime value
	port := env["PORT"]
	...
}

// defaultEnv returns the default values of this binary's settings
func defaultEnv() env.Builder {
	return env.Values{
		"PORT":":8080",
	}
}
```

This slightly more complex example uses `ParseDefault`, but maps 2 child `env.Values` to 2 `env`-aware example modules

```go
package main

import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)

func main() {
	// combine env defaults using prefixes
	env := env.New().Merge("DB_", db.DefaultEnv()).Merge("PAY_", pay.DefaultEnv())
	// parse runtime values
	env.ParseDefault() // ignoring error
	// produce db specific env k/v
	dbEnv := env.Match("DB_") // extract env for keys beginning with "DB_"
	db := db.NewSystem(dbEnv) // pass it on
	// produce pay specific env k/v
	payEnv := env.Match("PAY_") // extract env for keys beginning with "PAY_"
	pay := pay.NewSystem(payEnv) // pass it on
	...
}
```

Another similar example using custom/multiple env files for the same example modules

```go
import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)

func main() {
	// db env stored in .envrc, do not respect CLI flags or os env
	dbEnv, _ := env.Builder(db.DefaultEnv()).ParseFile(".envrc") // ignoring error
	db := db.NewSystem(dbEnv) // pass it on
	// pay env stored in custom file, respects flags and os ENV
	payEnv, _ := env.Builder(pay.DefaultEnv()).ParseFile(".paysys") // ignoring error
	payEnv.ParseShell().ParseDefaultArgs() // apply shell context, then CLI
	pay := pay.NewSystem(payEnv) // pass it on
	// rebuild sys env
	sysEnv := env.New().Merge("db.", dbEnv).Merge("pay.", payEnv)
	...
}
```

## `cmd/dotenv`

Simple binary that echos the default env behavior

```sh
$ go get taylz.io/env/cmd/dotenv
...
$ dotenv
dotenv: open .env: The system cannot find the file specified.
dotenv: env is empty
$ dotenv version
taylz.io/env/cmd/dotenv@v0.1.1
$ dotenv -hello
hello=true
$ touch .env
$ dotenv
dotenv: env is empty
$ echo ENV=pro > .env
$ dotenv
ENV=pro
$ dotenv -a ENV=dev -flag=x flag=y
ENV=dev
a=true
flag=y
```

## Changes

- v0.1.1
  - fix README
  - add `cmd.Parse1`
    - consumes `os.Args`, returns the first term as a "command name", and parse the remaining args with `ParseArgs`
- v0.1.0
    - `Values` is now an alias type, meaning you can easily consume `map[string]string` instead of `env.Values`
