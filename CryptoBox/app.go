package main

import (
	"context"
	"cryptolab/codebox"
	"encoding/base64"
	"encoding/hex"
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

func (a *App) Ed25519Sign(PrivateKeyStr string, PublicKeyStr string, msg string, format string) *codebox.SignResult {
	GoPrivateKeyStr := PrivateKeyStr + PublicKeyStr
	SignByte, err := codebox.Ed25519Sign(GoPrivateKeyStr, msg)

	var signature string

	if format == "base64" {
		signature = base64.StdEncoding.EncodeToString(SignByte)
	} else {
		signature = hex.EncodeToString(SignByte)
	}

	resp := &codebox.SignResult{
		Signature: signature,
	}

	if err != nil {
		resp.ErrMsg = err.Error()
	}

	return resp
}

func (a *App) Ed25519Verify(PubkeyStr string, msg string, signature string) *codebox.SignVerifyResult {
	isvalid, err := codebox.Ed25519Verify(PubkeyStr, msg, signature)

	resp := &codebox.SignVerifyResult{
		IsValid: isvalid,
	}

	if err != nil {
		resp.ErrMsg = err.Error()
	}

	return resp
}
