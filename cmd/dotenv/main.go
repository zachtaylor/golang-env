// Package dotenv is an executable that prints all values in the global env
package main

import (
	"fmt"
	"os"

	"taylz.io/env"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("taylz.io/env/cmd/dotenv@" + env.Version)
	} else if env, err := env.ParseDefault(); err != nil {
		fmt.Println("dotenv:", err.Error())
	} else if len(env) < 1 {
		fmt.Println("dotenv: env is empty")
	} else {
		for k, v := range env {
			fmt.Printf(k + "=" + v + "\n")
		}
	}
}
