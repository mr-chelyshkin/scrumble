//go:build race
// +build race

package version

// isRace reports whether the current binary was built with the Go
// race detector enabled.
const isRace = true
