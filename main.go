package main

import (
	"embed"
	"flag"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cliMode := flag.Bool("cli", false, "Run in CLI mode")
	flag.Parse()

	if *cliMode || len(os.Args) > 1 && os.Args[1] != "--cli" {
		runCLI()
		return
	}

	runGUI()
}

func runCLI() {
	println("================================================")
	println("Welcome to Password Manager")
	println("================================================")
	println("This is a CLI mode")
	println("================================================")
}

func runGUI() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "local-pwd-manager",
		Width:  1024,
		Height: 768,
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
