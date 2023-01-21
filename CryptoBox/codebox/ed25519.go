package codebox

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"

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

	// fmt.Println("privKeyBuf: ", privKeyBuf)
	// fmt.Println("pubBuf: ", pubBuf)

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

func Ed25519Sign(PrivateKeyStr string, msg string) ([]byte, error) {
	var privb []byte
	var err error

	// fmt.Println(">>>>: ", PrivateKeyStr)

	if len(PrivateKeyStr) == 128 {
		privb, err = hex.DecodeString(PrivateKeyStr)
	} else {
		privb, err = base64.StdEncoding.DecodeString(PrivateKeyStr)
	}

	if err != nil {
		return []byte{}, err
	}

	if len(privb) != 64 {
		return nil, errors.New("invalid private key size")
	}

	result := ed25519.Sign(privb, []byte(msg))

	return result, nil
}

func Ed25519Verify(PubKeyStr string, msg string, signatureStr string) (bool, error) {
	msgByte := []byte(msg)

	if PubKeyStr == "" {
		return false, errors.New("invalid pubkey")
	}

	var PublicKey []byte
	var err error

	if len(PubKeyStr) == 64 {
		PublicKey, err = hex.DecodeString(PubKeyStr)
		if err != nil {
			return false, err
		}
	} else {
		PublicKey, err = base64.StdEncoding.DecodeString(PubKeyStr)
		if err != nil {
			return false, err
		}
	}

	signature, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return false, err
	}

	isValid := ed25519.Verify(PublicKey, msgByte, signature)

	return isValid, nil
}
