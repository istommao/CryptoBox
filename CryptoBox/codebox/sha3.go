package codebox

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func CreateSHA3Hash(plainText string) *SHA3Result {
	sha224 := fmt.Sprintf("%x", sha3.Sum224([]byte(plainText)))
	sha256 := fmt.Sprintf("%x", sha3.Sum256([]byte(plainText)))
	sha384 := fmt.Sprintf("%x", sha3.Sum384([]byte(plainText)))
	sha512 := fmt.Sprintf("%x", sha3.Sum512([]byte(plainText)))

	result := &SHA3Result{}
	result.SHA3_224 = sha224
	result.SHA3_256 = sha256
	result.SHA3_384 = sha384
	result.SHA3_512 = sha512
	return result
}
