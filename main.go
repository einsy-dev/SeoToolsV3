package main

import (
	"embed"
	"fmt"

	"github.com/einsy-dev/NetSail/internal/services"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	app := application.New(application.Options{
		Name: "NetSail",
		Services: []application.Service{
			application.NewService(&services.Clip{}),
			application.NewService(&services.Knot{}),
			application.NewService(&services.Player{}),
			application.NewService(&services.Api{}),
			application.NewService(&services.CsvParser{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:     "NetSail",
		X:         1000,
		Y:         500,
		Width:     350,
		MinWidth:  350,
		MaxWidth:  700,
		Height:    500,
		MinHeight: 500,
		MaxHeight: 800})

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
