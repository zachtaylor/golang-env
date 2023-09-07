package env

import (
	"runtime/debug"
	"time"
)

var BuildInfo struct {
	Version  string
	Revision string
	Time     time.Time
	Modified bool
}

func init() {
	BuildInfo.Version = "(dev)"
	BuildInfo.Revision = `unknown`
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	if info.Main.Version != "" {
		BuildInfo.Version = info.Main.Version
	}
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			BuildInfo.Revision = kv.Value
		case "vcs.time":
			BuildInfo.Time, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			BuildInfo.Modified = kv.Value == "true"
		}
	}
}
