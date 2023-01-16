package codebox

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"fmt"
	"os"
)

func CreateEd25519KeyPair(format string) (string, string) {
	pub, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	privKeyBuf := make([]byte, len(privKey))
	copy(privKeyBuf, privKey)

	pubBuf := make([]byte, len(pub))
	copy(pubBuf, pub)

	var PrivateKey, PubKey string
	if format == "hex" {
		PrivateKey = hex.EncodeToString(privKeyBuf[:32])
		PubKey = hex.EncodeToString(pubBuf)
	} else {
		PrivateKey = base64.StdEncoding.EncodeToString(privKeyBuf[:32])
		PubKey = base64.StdEncoding.EncodeToString(pubBuf)
	}

	return PrivateKey, PubKey
}
