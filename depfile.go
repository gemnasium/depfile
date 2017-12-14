package depfile

type DepFile struct {
	Name            string           `json:"name"`                // used as a unique key
	Filenames       []string         `json:"filenames,omitempty"` // supported filenames
	Patterns        []string         `json:"patterns,omitempty"`  // patterns of supported filenames
	URL             string           `json:"url,omitempty"`       // documentation URL if any
	PackageManagers []PackageManager `json:"package_managers"`    // package managers capable of processing the file
}

type PackageManager struct {
	Name string `json:"name"`          // human readable name
	URL  string `json:"url,omitempty"` // documentation URL
}
