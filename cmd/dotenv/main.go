// Package dotenv is an executable that prints all values in the default env
package main

import (
	"fmt"
	"os"

	"taylz.io/env"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("taylz.io/env/cmd/dotenv@" + env.BuildInfo.Version + " [" + env.BuildInfo.Revision + "]")
		return
	}
	env, err := env.ParseDefault()
	if err != nil {
		fmt.Println("dotenv:", err.Error())
	}
	if len(env) < 1 {
		fmt.Println("dotenv: env is empty")
	} else {
		for k, v := range env {
			fmt.Println(k + "=" + v)
		}
	}
}
