package main

import (
	"context"
	"monsi/util"
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

func (a *App) GetListOfDIDs() []util.DID {
	return wallet.GetDIDs()
}

func (a *App) GetListOfVCs(did string) []string {
	return vcmanager.GetVCsAsStrings(did)
}

func (a *App) RefreshVCs() {
	vcmanager.RefreshVCs()
}

func (a *App) StoreVC(vcName string, vcContent string) {
	vcmanager.StoreVC(vcName, vcContent)
}
