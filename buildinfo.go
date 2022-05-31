package buildinfo

import (
	"os/exec"
	"runtime/debug"
	"strings"
)

type Info struct {
	GoVersion string
	GitCommit string
	GitTag    string
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

			cmd := exec.Command("git", "describe", "--tags", info.GitCommit)
			cmdResult, err := cmd.Output()
			if err != nil {
				// Swallow error
				return info, true
			}

			info.GitTag = strings.TrimSpace(string(cmdResult))
		}
	}

	return info, true
}
