# taylz.io/env

Package env provides runtime environment, using flags, file, and/or environment variables

## `env.Values`

Centralized runtime values with defaults

```go
import "taylz.io/env"
...
var env = env.Values{
	...
	"PORT":":8080",
	...
}
...
func main() {
	env.ParseDefault() // parse ".env" file, os env, and cli args
	...
	port := env["PORT"] // access runtime value
}
```

### Environment management

Using common configuration

```go
import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)
...
func main() {
	// combine env defaults using prefix
	env := env.New().Merge("DB_", db.DefaultSettings()).Merge("PAY_", pay.DefaultSettings())
	// parse runtime values
	env.ParseDefault()
	...
	// db env
	dbEnv := env.Match("DB_") // extact env for keys beginning with "DB_"
	dbConn := internal.NewDBConn(dbEnv) // client uses env
	...
	// pay env
	payEnv := env.Match("PAY_") // extract env for keys beginning with "PAY_"
	payService := internal.NewPayService(payEnv) // client uses env
	...
}
```

Using seperate configuration

```go
import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)
...
func main() {
	// db env
	dbEnv := db.DefaultSettings().ParseFile(".db.env")
	dbConn := internal.NewDBConn(dbEnv) // client uses env
	// pay env
	payEnv := pay.DefaultSettings().ParseFile(".pay.env")
	payService := internal.NewPayService(payEnv) // client uses env
	...
}
```

## File format

The file format is basic shell-like format

```sh
# .env file
PORT=:8080
```

## CLI format

The command line argument format is `k=v`

For `k` starting with `-`, the leading `-` is removed

For `v` ending with `\`, the following term is concatenated to the value of `v`, following a space

# cmd/dotenv

Simple binary that echos the default env behavior

```sh
$ go get taylz.io/cmd/dotenv
...
$ dotenv
dotenv: open .env: The system cannot find the file specified.
dotenv: env is empty
$ dotenv version
dotenv version v0.0.0
$ dotenv -hello
dotenv: open .env: The system cannot find the file specified.
hello=true
$ touch .env
$ dotenv
env is empty
$ echo ENV=pro > .env
$ dotenv
ENV=pro
$ dotenv -a ENV=dev -flag=x flag=y
ENV=dev
a=true
flag=y
```
