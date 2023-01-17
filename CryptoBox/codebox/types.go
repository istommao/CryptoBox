package codebox

type KeyPair struct {
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`
}

type SignResult struct {
	Signature string `json:"signature"`
	ErrMsg    string `json:"errmsg"`
}
