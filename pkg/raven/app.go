package raven

import (
	"context"
	"fmt"

	"github.com/soikes/raven/pkg/appdata"
	"github.com/soikes/raven/pkg/browser"
	"github.com/soikes/raven/pkg/db"
	"github.com/soikes/raven/pkg/fetch/firefox"
	"github.com/soikes/raven/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appData *appdata.Config
	db      *db.Db
	fetcher *firefox.Config
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
}

func (a *App) Fetch() {}
func (a *App) Run()   {}
func (a *App) FetchVersions() ([]types.BrowserListItem, error) {
	versions, err := firefox.GetVersions()
	if err != nil {
		return nil, err
	}
	var bli []types.BrowserListItem
	for _, version := range versions {
		b := browser.Browser{
			Version:   version,
			Name:      `Firefox`,
			Path:      ``,
			Available: false,
		}
		// TODO: Fallback to the DB (Or start with the DB?)
		// TODO: Where to generate the path
		// TODO: Check the filesystem to see if it is installed
		// TODO: Consolidate all of the above in a proper place ("sync")
		bli = append(bli, types.NewBrowserListItem(b))
	}
	return bli, nil
}

func failStart(ctx context.Context, err error) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Message: fmt.Sprintf("Raven was unable to start:\n%s", err.Error()),
		Buttons: []string{`Ok`},
	})
}
