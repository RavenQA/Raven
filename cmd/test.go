package main

import (
	"os"

	"github.com/soikes/raven/pkg/installer/dmg"
	"github.com/soikes/raven/pkg/run"
)

func main() {
	mnt, err := os.MkdirTemp("", "RavenBrowserVolumes")
	if err != nil {
		panic(err)
	}
	d := dmg.NewDmg(dmg.WithMountPoint(mnt))
	app := "/Users/michael/Local/Firefox.app"
	err = d.ExtractApp("/Users/michael/Downloads/Firefox 135.0.dmg", app)
	if err != nil {
		panic(err)
	}
	err = run.RunMacOS(app, "-new-window", "http://soikke.li/")
	if err != nil {
		panic(err)
	}
}

// cfg := firefox.Config{
// 	Platform: fetch.PlatformMac,
// }
// err := cfg.Fetch(`latest`, language.AmericanEnglish)
// if err != nil {
// 	panic(err)
// }
