package common

// Version is set during build via ldflags
var Version = "dev"

// GetVersion returns the current version of the application
func GetVersion() string {
	return Version
}
