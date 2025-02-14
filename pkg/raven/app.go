package raven

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"slices"

	"github.com/soikes/raven/pkg/appdata"
	"github.com/soikes/raven/pkg/browser"
	"github.com/soikes/raven/pkg/browser/firefox"
	"github.com/soikes/raven/pkg/browser/firefox/fetch"
	"github.com/soikes/raven/pkg/db"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appData *appdata.Config
	db      *db.Db
	ff      firefox.Firefox
	ctx     context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Start saves the runtime context and initializes all required
// backend services for the app to function.
// If Start() fails, an error dialog is shown to the user then the app panics.
func (a *App) Start(ctx context.Context) {
	a.ctx = ctx
	data, err := appdata.NewConfig()
	if err != nil {
		failStart(ctx, err)
		panic(err)
	}
	a.appData = data
	d, err := db.NewDB(a.appData.Dir)
	if err != nil {
		failStart(ctx, err)
		panic(err)
	}
	a.db = d
	err = a.db.Init(ctx)
	if err != nil {
		failStart(ctx, err)
		panic(err)
	}
	a.ff = firefox.Firefox{Ctx: ctx, Db: d}
}

func (a *App) SyncBrowsers() ([]browser.Browser, error) {
	browsers, err := a.db.GetBrowsers(a.ctx)
	if err != nil {
		return nil, err
	}
	remoteBrowsers, err := fetch.FetchBrowserList()
	if err != nil {
		log.Printf("failed to fetch new browsers: %s", err.Error())
		if len(browsers) == 0 {
			return nil, errors.New("Failed to fetch browser list. Check your network connection and try again later.")
		} else {
			return browsers, nil
		}
	}
	var newBrowsers []browser.Browser
	for _, remote := range remoteBrowsers {
		newBrowser := true
		for _, local := range browsers {
			if remote.Identifier() == local.Identifier() {
				newBrowser = false
			}
		}
		if newBrowser {
			newBrowsers = append(newBrowsers, remote)
		}
	}
	browsers = append(browsers, newBrowsers...)
	err = a.db.InsertBrowsers(a.ctx, browsers)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(browsers, func(a, b browser.Browser) int {
		if a.ReleaseDate.Equal(b.ReleaseDate) {
			return 0
		}
		if a.ReleaseDate.Before(b.ReleaseDate) {
			return -1
		}
		return 1
	})
	slices.Reverse(browsers)
	return browsers, nil
}

func (a *App) FetchFirefox(version string) error {
	id := browser.Identifier(browser.ProductFirefox, version)
	cfg := browser.FetchConfig{
		Version:      version,
		TmpDir:       a.appData.TmpDir,
		DownloadName: id,
	}
	err := a.ff.Fetch(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) InstallFirefox(version string) error {
	id := browser.Identifier(browser.ProductFirefox, version)
	installPath := filepath.Join(a.appData.Dir, id)
	icfg := browser.InstallConfig{
		ImageName:  id,
		AppPath:    installPath,
		TmpDir:     a.appData.TmpDir,
		VolumesDir: a.appData.TmpDir,
	}
	err := a.ff.Install(icfg)
	if err != nil {
		return err
	}
	err = a.db.UpdateInstallPath(a.ctx, installPath, browser.ProductFirefox, version)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) LaunchFirefox(version, startUrl string) error {
	id := browser.Identifier(browser.ProductFirefox, version)
	cfg := browser.LaunchConfig{
		AppPath:  filepath.Join(a.appData.Dir, id),
		StartUrl: startUrl,
	}
	return a.ff.Launch(cfg)
}

func (a *App) DropSchema() error {
	return a.db.DropAll(a.ctx)
}

func failStart(ctx context.Context, err error) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Message: fmt.Sprintf("Raven was unable to start:\n%s", err.Error()),
		Buttons: []string{`Ok`},
	})
}
