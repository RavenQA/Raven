package browser

type FetchConfig struct {
	TmpDir string
}

type InstallConfig struct {
	ImagePath  string
	VolumesDir string
	AppDir     string
}

type LaunchConfig struct {
	AppDir   string
	StartUrl string
}
