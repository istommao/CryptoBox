package codebox

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func CreateSHA1MD5Hash(plainText string) *SHA1MD5Result {

	sha1Result := fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	MD5Result := fmt.Sprintf("%x", md5.Sum([]byte(plainText)))

	result := &SHA1MD5Result{}
	result.SHA1 = sha1Result
	result.MD5 = MD5Result

	return result
}
