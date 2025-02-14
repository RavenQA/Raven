package browser

import "github.com/soikes/raven/pkg/browser/firefox/policy"

type FetchConfig struct {
	Version      string
	TmpDir       string
	DownloadName string
}

type InstallConfig struct {
	TmpDir     string
	ImageName  string
	VolumesDir string
	AppPath    string
}

type LaunchConfig struct {
	AppPath  string
	StartUrl string
	Policy   *policy.PolicyRoot
}
