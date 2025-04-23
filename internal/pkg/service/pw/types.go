package pw

type PasswordPair struct {
	EncryptedHash string `json:"encrypted_hash"`
	Nonce         string `json:"nonce"`
}
