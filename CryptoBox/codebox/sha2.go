package codebox

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func CreateSHA2Hash(plainText string) *SHA2Result {

	sha224 := fmt.Sprintf("%x", sha256.Sum224([]byte(plainText)))
	sha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(plainText)))
	sha384 := fmt.Sprintf("%x", sha512.Sum384([]byte(plainText)))
	sha512 := fmt.Sprintf("%x", sha512.Sum512([]byte(plainText)))

	result := &SHA2Result{}
	result.SHA224 = sha224
	result.SHA256 = sha256
	result.SHA384 = sha384
	result.SHA512 = sha512
	return result
}
