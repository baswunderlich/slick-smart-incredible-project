package main

import (
	"embed"
	"fmt"
	"monsi/api"
	"monsi/vcmanager"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	vcmanager.ReadVCsFromFiles()
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
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
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
	router.POST("api/vc/sign", api.SignVC)
	router.GET("api/did", api.ListDIDs)
	router.POST("api/encrypt", api.Encrypt)
	router.POST("api/decrypt", api.Decrypt)
	router.POST("api/mail", api.RecieveMail)
	router.POST("api/mail/new", api.GenMail)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	router.Run("0.0.0.0:80")
}
