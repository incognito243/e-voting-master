package blockchain

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func GenRandomAptosAddress() (string, error) {
	randomPubKey := make([]byte, 32)
	_, err := rand.Read(randomPubKey)
	if err != nil {
		return "", err
	}

	hasher := sha3.New256()
	hasher.Write(randomPubKey)
	addressBytes := hasher.Sum(nil)

	address := "0x" + hex.EncodeToString(addressBytes)
	return address, nil
}
