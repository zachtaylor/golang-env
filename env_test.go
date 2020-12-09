package env_test

import (
	"fmt"
	"testing"

	"taylz.io/env"
)

func TestParseArgsSpace(t *testing.T) {
	env := env.ParseArgs([]string{
		"-z=Hello\\",
		"World!",
	})
	if env["z"] != "Hello World!" {
		t.Log("Expected: \"Hello World!\"")
		t.Log("Actual: \"" + env["z"] + "\"")
		t.Log(fmt.Sprint(env))
		t.Fail()
	}
}
