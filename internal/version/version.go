package version

import (
	"fmt"
	"runtime"
	"strings"
)

// Base version information.
//
// This is the fallback data used when version information from git.
var (
	gitVersion   = "v0.0.0-main"
	gitCommit    = "unknown" // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "unknown" // state of git tree, either "clean" or "dirty"

	buildDate = "unknown" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Info contains versioning information.
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
	Race         bool   `json:"race"`
}

// String returns a Go-syntax representation of the Info.
func (i Info) String() string {
	return fmt.Sprintf("%#v", i)
}

func (i Info) Pretty() string {
	var b strings.Builder
	b.WriteString(i.GoVersion)
	b.WriteRune(' ')
	b.WriteString(i.Platform)
	if i.Race {
		b.WriteRune(' ')
		b.WriteString("(race detector enabled)")
	}

	return b.String()
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Race:         isRace,
	}
}
