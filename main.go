package main

import (
	"embed"

	"github.com/einsy-dev/SEOTools/internal/services"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name: "Seo Tools",
		Services: []application.Service{
			application.NewService(&services.CSV{}),
			application.NewService(&services.Clipboard{}),
			application.NewService(&services.Player{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:       "SEO Tools",
		X:           1000,
		Y:           500,
		Width:       350,
		MinWidth:    350,
		Height:      500,
		MinHeight:   500,
		AlwaysOnTop: true,
	})

	err := app.Run()
	if err != nil {
		panic(err)
	}

	// err := wails.Run(&options.App{
	// 	Title:       "SEO Tools",
	// 	Width:       350,
	// 	MinWidth:    350,
	// 	Height:      500,
	// 	MinHeight:   500,
	// 	AlwaysOnTop: true,
	// 	AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
	// 	OnStartup: func(ctx context.Context) {
	// 		app.Ctx = ctx
	// 		clipboard.Startup()
	// 		extension.Startup()
	// 		sql.Sql.Startup()
	// 	},
	// 	Bind: []interface{}{
	// 		extension.Bind,
	// 	},
	// })

	// if err != nil {
	// 	println("Error:", err.Error())
	// }
}
