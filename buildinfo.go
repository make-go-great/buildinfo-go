package buildinfo

import (
	"runtime/debug"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

			// Find tag from commit
			// If err, swallow as maybe this is not a git repo
			repo, err := git.PlainOpen(".")
			if err != nil {
				return info, true
			}

			hash, err := repo.ResolveRevision(plumbing.Revision(info.GitCommit))
			if err != nil {
				return info, true
			}

			tag, err := repo.TagObject(*hash)
			if err != nil {
				return info, true
			}

			info.GitTag = tag.Name
		}
	}

	return info, true
}
