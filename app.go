package main

import (
	"context"
	"os"

	"github.com/soikes/raven/pkg/fetch/firefox"
	"github.com/soikes/raven/pkg/installer/dmg"
	"github.com/soikes/raven/pkg/platform"
	"github.com/soikes/raven/pkg/run"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/language"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Fetch() error {
	fcfg := firefox.Config{
		Platform:     platform.PlatformMac,
		ProgressFunc: a.onProgress,
	}
	err := fcfg.Fetch(`latest`, language.AmericanEnglish)
	if err != nil {
		return err
	}
	return nil
}

func (a App) onProgress(progress, total int) {
	pct := float64(progress) / float64(total) * 100
	runtime.EventsEmit(a.ctx, "fetchProgress", pct)
}

func (a *App) Run() error {
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
	return nil
}
