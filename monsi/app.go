package main

import (
	"context"
	"monsi/wallet"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetListOfDIDs() string {
	dids := wallet.GetDIDs()
	res := ""
	for _, d := range dids {
		res += (d + "<br></br>")
	}
	return res
}

// Greet returns a greeting for the given name
func (a *App) AddDID() string {
	return "Hello! When everything is set up you should now be able to add a DID"
}
