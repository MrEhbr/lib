package version

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// These values are private which ensures they can only be set with the build flags.
var (
	version   = "unknown"
	branch    = "unknown"
	revision  = "unknown"
	goVersion = "unknown"
	buildDate = "unknown"
	buildUser = "unknown"
	appName   = "unknown"
)

// Info is a structure with version build information about the current application.
type Info struct {
	Version   string `json:"version"`
	Branch    string `json:"branch"`
	Revision  string `json:"revision"`
	GoVersion string `json:"go_version"`
	BuildDate string `json:"build_date"`
	BuildUser string `json:"build_user"`
}

// Version returns a structure with the current version information.
func Version() Info {
	return Info{
		Version:   version,
		Branch:    branch,
		Revision:  revision,
		GoVersion: goVersion,
		BuildDate: buildDate,
		BuildUser: buildUser,
	}
}

// Print outputs the application name and version string.
func Print() {
	v := Version()
	fmt.Printf("%s version %s\n", appName, v.Version)
}

// PrintFull prints the application name and detailed version information.
func PrintFull() {
	v := Version()
	fmt.Printf("%s - version %s\n", appName, v.Version)
	fmt.Printf("  branch: \t%s\n", v.Branch)
	fmt.Printf("  revision: \t%s\n", v.Revision)
	fmt.Printf("  build date: \t%s\n", v.BuildDate)
	fmt.Printf("  build user: \t%s\n", v.BuildUser)
	fmt.Printf("  go version: \t%s\n", v.GoVersion)
}

// Handler returns an HTTP Handler which returns JSON formatted version information.
func Handler() http.Handler {
	v := Version()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		_ = enc.Encode(v)
	})
}
