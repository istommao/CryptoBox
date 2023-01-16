package main

import (
	"context"
	"cryptolab/codebox"
	"fmt"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateEd25519KeyPair(format string) *codebox.KeyPair {
	privateKey, pubkey := codebox.CreateEd25519KeyPair(format)
	// fmt.Println(privateKey, pubkey)

	keyPair := &codebox.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  pubkey,
	}
	return keyPair
}
