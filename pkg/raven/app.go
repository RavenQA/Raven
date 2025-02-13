package raven

import (
	"context"
	"fmt"

	"github.com/soikes/raven/pkg/appdata"
	"github.com/soikes/raven/pkg/db"
	"github.com/soikes/raven/pkg/fetch/firefox"
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

// func (a *App) GetBrowsers() []browser.Browser {

// }

func failStart(ctx context.Context, err error) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Message: fmt.Sprintf("Raven was unable to start:\n%s", err.Error()),
		Buttons: []string{`Ok`},
	})
}
