package browser

import (
	"os"

	"github.com/soikes/raven/pkg/installer/dmg"
	"github.com/soikes/raven/pkg/run"
)

type Firefox struct{}

func (f *Firefox) Launch() error {
	mnt, err := os.MkdirTemp("", "RavenBrowserVolumes")
	if err != nil {
		return err
	}
	d := dmg.NewDmg(dmg.WithMountPoint(mnt))
	app := "/Users/michael/Local/Firefox.app"
	err = d.ExtractApp("/Users/michael/Downloads/Firefox 135.0.dmg", app)
	if err != nil {
		return err
	}
	err = run.RunMacOS(app, "-new-window", "http://soikke.li/")
	if err != nil {
		return err
	}
	return nil
}
