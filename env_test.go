package env_test

import (
	"testing"

	"taylz.io/env"
)

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
