# taylz.io/env

Package env provides runtime environment, using flags, file, and/or environment variables

## Example `main.go`

```go
import "taylz.io/env"

func main() {
	env, _ := defaultEnv().ParseDefault() // parse ".env" file, os env, and cli args
	...
	port := env["PORT"] // access runtime value
}

func defaultEnv() env.Values {
	return env.Values{
		env.Values{
		...
		"PORT":":8080",
		...
	}
}
```

## Sub env examples

Using combined env

```go
import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)

func main() {
	// combine env defaults using prefixes
	env := env.New().Merge("DB_", db.DefaultEnv()).Merge("PAY_", pay.DefaultEnv())
	// parse runtime values
	env.ParseDefault()
	...
	// db env
	dbEnv := env.Match("DB_") // extract env for keys beginning with "DB_"
	dbConn := internal.NewDBConn(dbEnv) // client uses env
	...
	// pay env
	payEnv := env.Match("PAY_") // extract env for keys beginning with "PAY_"
	payService := internal.NewPayService(payEnv) // client uses env
	...
}
```

Using seperate/multiple env files

```go
import (
	"example.com/db"
	"example.com/pay"
	"taylz.io/env"
)

func main() {
	// db env
	dbEnv := db.DefaultEnv().ParseFile(".db.env")
	dbConn := internal.NewDBConn(dbEnv) // client uses env
	...
	// pay env
	payEnv := pay.DefaultEnv().ParseFile(".pay.env")
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

## Parse format

Keys should not begin with `-`, as `Parse` will remove all leading `-`

### `ParseArgs` space format ("")

Golang has default support for traditional "quotation based grouping"

## `cmd/dotenv`

Simple binary that echos the default env behavior

```sh
$ go get taylz.io/cmd/dotenv
...
$ dotenv
dotenv: open .env: The system cannot find the file specified.
dotenv: env is empty
$ dotenv version
dotenv version v0.0.1
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
