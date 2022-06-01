package buildinfo

import (
	"runtime/debug"
)

type Info struct {
	GoVersion string

	// https://github.com/golang/go/issues/29228
	// https://github.com/golang/go/issues/50603
	MainModuleVersion string

	GitCommit string
}

func Read() (*Info, bool) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, false
	}

	info := &Info{
		GoVersion:         buildInfo.GoVersion,
		MainModuleVersion: buildInfo.Main.Version,
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
