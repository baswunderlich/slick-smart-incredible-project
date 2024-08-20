package main

import (
	"embed"
	"fmt"
	"monsi/api"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go startRouter()
	startWails()
}

func startWails() {
	fmt.Println("Wails started ...")
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "monsi",
		Width:  800,
		Height: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func startRouter() {
	fmt.Println("Router started ...")
	router := gin.Default()
	router.POST("api/vc", api.ListVCs)
	router.GET("api/did", api.ListDIDs)

	router.Run("0.0.0.0:80")
}
