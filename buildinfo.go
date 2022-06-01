package buildinfo

import (
	"runtime/debug"
)

type Info struct {
	GoVersion string
	GitCommit string
}

func Read() (*Info, bool) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, false
	}

	info := &Info{
		GoVersion: buildInfo.GoVersion,
	}

	for _, kv := range buildInfo.Settings {
		switch kv.Key {
		case "vcs.revision":
			// Assume version control is git
			info.GitCommit = kv.Value
		}
	}

	return info, true
}
