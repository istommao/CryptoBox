package codebox

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/curve25519"
)

func CreateX25519KeyPair() (string, string) {
	rand.Seed(time.Now().UnixNano())

	var privateKey [32]byte
	for i := range privateKey[:] {
		privateKey[i] = byte(rand.Intn(256))
	}

	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)

	PrivateKeyStr := fmt.Sprintf("%x", privateKey)

	PubKeyStr := fmt.Sprintf("%x", publicKey)

	return PrivateKeyStr, PubKeyStr
}
