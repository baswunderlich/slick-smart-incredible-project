package main

import (
	"context"
	"monsi/vcmanager"
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

func (a *App) GetListOfDIDs() []string {
	return wallet.GetDIDs()
}

func (a *App) GetListOfVCs(did string) []string {
	return vcmanager.GetVCs(did)
}

// Greet returns a greeting for the given name
func (a *App) AddDID() string {
	return "Hello! When everything is set up you should now be able to add a DID"
}
