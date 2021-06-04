package env_test

import (
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
		t.Log(env)
		t.Fail()
	}
}

func TestParseArgsHyphen(t *testing.T) {
	env := env.ParseArgs([]string{"a", "-b", "--c"})

	if env["a"] != "true" {
		t.Log(`Expected: a="true"`)
		t.Log(`Actual: "` + env["a"] + `"`)
		t.Log(env)
		t.Fail()
	}
	if env["b"] != "true" {
		t.Log(`Expected: b="true"`)
		t.Log(`Actual: "` + env["b"] + `"`)
		t.Log(env)
		t.Fail()
	}
	if env["c"] != "true" {
		t.Log(`Expected: c="true"`)
		t.Log(`Actual: "` + env["c"] + `"`)
		t.Log(env)
		t.Fail()
	}
}
