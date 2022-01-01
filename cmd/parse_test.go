package cmd_test

import (
	"os"
	"testing"

	"taylz.io/env/cmd"
)

func TestParseCmd1(t *testing.T) {
	oldargs := os.Args

	os.Args = []string{}
	cmd1, env1 := cmd.Parse1(nil)
	if len(cmd1) > 0 || env1 != nil {
		t.Log(`Expected safe from bamboozle`)
		t.Log(`Actual: "`+cmd1+`", `, env1)
		t.Fail()
	}

	os.Args = []string{"example.exe", "version"}
	cmd2, env2 := cmd.Parse1(nil)
	if cmd2 != "version" || env2 == nil {
		t.Log(`Expected cmd="version" env={}`)
		t.Log(`Actual: "`+cmd2+`", `, env2)
		t.Fail()
	}

	os.Args = []string{"git", "clone", "example.com/fubar"}
	cmd3, env3 := cmd.Parse1(nil)
	if cmd3 != "clone" {
		t.Log(`Expected cmd="clone"`)
		t.Log(`Actual: "` + cmd2 + `"`)
		t.Fail()
	}
	if len(env3) != 1 || env3["example.com/fubar"] != "true" {
		t.Log(`Expected env["example.com/fubar"]="true"`)
		t.Log(`Actual: "` + env3["example.com/fubar"] + `"`)
		t.Fail()
	}

	os.Args = oldargs
}
