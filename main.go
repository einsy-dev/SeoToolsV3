package main

import (
	"context"
	"embed"

	"github.com/einsy-dev/SEOTools/internal/app"
	"github.com/einsy-dev/SEOTools/internal/clipboard"
	"github.com/einsy-dev/SEOTools/internal/extension"
	"github.com/einsy-dev/SEOTools/internal/sql"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	err := wails.Run(&options.App{
		Title:       "SEO Tools",
		Width:       350,
		MinWidth:    350,
		Height:      500,
		MinHeight:   500,
		AlwaysOnTop: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Ctx = ctx
			clipboard.Startup()
			extension.Startup()
			sql.Sql.Startup()
		},
		Bind: []interface{}{
			extension.Bind,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
