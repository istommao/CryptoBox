package codebox

type KeyPair struct {
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`
}

type SignResult struct {
	Signature string `json:"signature"`
	ErrMsg    string `json:"errmsg"`
}

type SignVerifyResult struct {
	IsValid bool   `json:"isvalid"`
	ErrMsg  string `json:"errmsg"`
}

type SHA1MD5Result struct {
	SHA1 string `json:"sha1"`
	MD5  string `json:"md5"`
}

type SHA2Result struct {
	SHA224 string `json:"sha224"`
	SHA256 string `json:"sha256"`
	SHA384 string `json:"sha384"`
	SHA512 string `json:"sha512"`
}

type SHA3Result struct {
	SHA3_224 string `json:"sha3_224"`
	SHA3_256 string `json:"sha3_256"`
	SHA3_384 string `json:"sha3_384"`
	SHA3_512 string `json:"sha3_512"`
}

type RSAKeyPair struct {
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`

	ErrMsg string `json:"errmsg"`
}

type AesResult struct {
	CipherText string `json:"CipherText"`
	Format     string `json:"format"`

	ErrMsg string `json:"errmsg"`
}
