package move

import (
	"errors"
	"os"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/crypto"
)

type Move struct {
	account *aptos.Account
}

func NewMove(executorAddress string) (*Move, error) {
	privateKeyHex := os.Getenv(executorAddress)
	if privateKeyHex == "" {
		return nil, errors.New("can not find environment variable")
	}

	address := &aptos.AccountAddress{}
	err := address.ParseStringRelaxed(executorAddress)
	if err != nil {
		panic("Failed to parse address:" + err.Error())
	}
	privateKey := &crypto.Ed25519PrivateKey{}
	err = privateKey.FromHex(privateKeyHex)
	if err != nil {
		panic("Failed to parse private key:" + err.Error())
	}

	addressHex, _ := address.MarshalJSON()
	account, err := aptos.NewAccountFromSigner(privateKey, addressHex)
	if err != nil {
		return nil, err
	}

	client, err := aptos.NewClient(aptos.NetworkConfig{})

	return &Move{
		account: account,
	}, nil
}
