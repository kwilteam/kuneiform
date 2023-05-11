package version

// Default build-time variable.
// These values are overridden via ldflags
// THese variables are used by goreleaser by default
var (
	Version = "unknown-version"
	Commit  = "unknown-commit"
	Date    = "unknown-buildtime"
)
