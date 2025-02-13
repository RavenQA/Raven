package firefox

import (
	"context"
	"os"
	"path/filepath"

	"github.com/soikes/raven/pkg/browser"
	"github.com/soikes/raven/pkg/fetch/firefox"
	"github.com/soikes/raven/pkg/installer/dmg"
	"github.com/soikes/raven/pkg/platform"
	"github.com/soikes/raven/pkg/progress"
	"github.com/soikes/raven/pkg/run"
	"golang.org/x/text/language"
)

type Firefox struct {
	browser.Browser
}

func (f *Firefox) Install(cfg browser.InstallConfig) error {
	mnt, err := os.MkdirTemp(cfg.VolumesDir, "Volume")
	if err != nil {
		return err
	}
	d := dmg.NewDmg(
		dmg.WithMountPoint(mnt),
		dmg.WithAppPath(filepath.Join(cfg.AppDir, f.Identifier())),
		dmg.WithImagePath(cfg.ImagePath),
	)
	err = d.Install()
	if err != nil {
		return err
	}
	return nil
}

func (f *Firefox) Launch(cfg browser.LaunchConfig) error {
	name := f.Identifier()
	err := run.RunMacOS(filepath.Join(cfg.AppDir, name), "-new-window", cfg.StartUrl)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firefox) Fetch(ctx context.Context, cfg browser.FetchConfig) error {
	fcfg := firefox.Config{
		Platform:     platform.PlatformMac,
		ProgressFunc: progress.ProgressPercentageHandler(ctx),
	}
	err := fcfg.Fetch(f.VersionString(), language.AmericanEnglish)
	if err != nil {
		return err
	}
	return nil
}
