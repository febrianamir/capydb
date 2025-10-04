package main

import (
	"capydb/backend/usecase"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	usecase := usecase.NewUsecase()
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CapyDB",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
			usecase,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
