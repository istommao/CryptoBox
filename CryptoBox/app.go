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

func (a *App) CreateSHA1MD5Hash(msg string) *codebox.SHA1MD5Result {
	return codebox.CreateSHA1MD5Hash(msg)
}

func (a *App) CreateSHA2Hash(msg string) *codebox.SHA2Result {
	return codebox.CreateSHA2Hash(msg)
}

func (a *App) CreateSHA3Hash(msg string) *codebox.SHA3Result {
	return codebox.CreateSHA3Hash(msg)
}

func (a *App) CreateRSAKeyPair(keySize int) *codebox.RSAKeyPair {
	resp := &codebox.RSAKeyPair{}

	prvKey, err := codebox.CreateRSAKeyPair(keySize)
	if err != nil {
		resp.ErrMsg = err.Error()
		return resp
	}

	PrivateKeyStr := codebox.ExportRsaPrivateKeyAsPemStr(prvKey)
	resp.PrivateKey = PrivateKeyStr

	PubkeyStr, err := codebox.ExportRsaPublicKeyAsPemStr(&prvKey.PublicKey)
	if err != nil {
		resp.ErrMsg = err.Error()
	} else {
		resp.PublicKey = PubkeyStr
	}

	return resp
}

func (a *App) CreateX25519KeyPair() *codebox.KeyPair {
	PrivateKey, PubKey := codebox.CreateX25519KeyPair()

	resp := &codebox.KeyPair{PrivateKey: PrivateKey, PublicKey: PubKey}

	return resp
}

func (a *App) AesGCMEncrypt(keyStr string, plainText string, nonceStr string) *codebox.AesResult {
	AesResultData := &codebox.AesResult{}

	key, err := hex.DecodeString(keyStr)
	if err != nil {
		AesResultData.ErrMsg = err.Error()

		return AesResultData
	}

	nonce, err := hex.DecodeString(nonceStr)
	if err != nil {
		AesResultData.ErrMsg = err.Error()

		return AesResultData
	}

	resultBytes, err := codebox.AesGCMEncrypt(key, plainText, nonce, nil)
	if err != nil {
		AesResultData.ErrMsg = err.Error()

		return AesResultData
	}

	AesResultData.Format = "hex"
	AesResultData.CipherText = hex.EncodeToString(resultBytes)

	return AesResultData
}
