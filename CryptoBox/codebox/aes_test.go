package codebox

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

const isBase64 = "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$"

func TestAesEncryptDecrypt(t *testing.T) {
	key := []byte("0123456789abcdef") // must be of 16 bytes for this example to work
	message := "Hello"

	encrypted, err := AesEncrypt(key, message)
	require.Nil(t, err)
	require.Regexp(t, isBase64, encrypted)

	decrypted, err := AesDecrypt(key, encrypted)
	require.Nil(t, err)
	require.Equal(t, message, decrypted)
}

func TestAesGCM(t *testing.T) {
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("Hello")

	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	ciphertext, err := AesGCMEncrypt(key, string(plaintext), nonce, nil)
	require.Nil(t, err)

	result, err := AesGCMDecrypt(key, ciphertext, nonce, nil)
	require.Nil(t, err)

	require.Equal(t, plaintext, result)
}
