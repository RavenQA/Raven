package firefox

import (
	"context"
	"os"
	"path/filepath"

	"github.com/soikes/raven/pkg/browser"
	"github.com/soikes/raven/pkg/browser/firefox/fetch"
	"github.com/soikes/raven/pkg/browser/firefox/policy"
	"github.com/soikes/raven/pkg/db"
	"github.com/soikes/raven/pkg/installer/dmg"
	"github.com/soikes/raven/pkg/platform"
	"github.com/soikes/raven/pkg/progress"
	"github.com/soikes/raven/pkg/run"
	"golang.org/x/text/language"
)

type Firefox struct {
	Db  *db.Db
	Ctx context.Context
}

func (f *Firefox) Fetch(cfg browser.FetchConfig) error {
	fcfg := fetch.Config{
		Platform:     platform.PlatformMac,
		ProgressFunc: progress.ProgressPercentageHandler(f.Ctx),
		DownloadPath: filepath.Join(cfg.TmpDir, cfg.DownloadName),
	}
	err := fcfg.Fetch(cfg.Version, language.AmericanEnglish)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firefox) Install(cfg browser.InstallConfig) error {
	mnt, err := os.MkdirTemp(cfg.VolumesDir, "Volume")
	if err != nil {
		return err
	}
	d := dmg.NewDmg(
		dmg.WithMountPoint(mnt),
		dmg.WithAppPath(cfg.AppPath),
		dmg.WithImagePath(filepath.Join(cfg.TmpDir, cfg.ImageName)),
	)
	err = d.Install(f.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firefox) Launch(cfg browser.LaunchConfig) error {
	var err error
	if cfg.Policy == nil {
		err = policy.DefaultPolicy.Save(cfg.AppPath)
	} else {
		err = cfg.Policy.Save(cfg.AppPath)
	}
	if err != nil {
		return err
	}
	err = run.RunMacOS(f.Ctx, cfg.AppPath, "-new-window", cfg.StartUrl)
	if err != nil {
		return err
	}
	return nil
}
