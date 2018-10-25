package version

import "fmt"

// Version info
var (
	Version   = "unknown"
	GitCommit = "unknown"
	BuildTime = "unknown"
)

// PrintVersion show version info on command line
func PrintVersion() {
	fmt.Printf(
		"Version: %s\nGit Commit: %s\nBuild Time: %s\n",
		Version,
		GitCommit,
		BuildTime,
	)
}
