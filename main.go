package main

import (
	"embed"

	"github.com/soikes/raven/pkg/browser"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	ff := browser.Firefox{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "raven",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&ff,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
