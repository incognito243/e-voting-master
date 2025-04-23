package blockchain

import (
	"errors"
	"fmt"
	"os"

	"e-voting-mater/configs"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/crypto"
)

type Client struct {
	client *aptos.Client
	sender *aptos.Account
}

func NewClient(account configs.VotingConfig, client *aptos.Client) (*Client, error) {
	key := os.Getenv(account.PrivateKeyName)
	privateKey := &crypto.Ed25519PrivateKey{}

	err := privateKey.FromHex(key)
	if err != nil {
		return nil, err
	}

	sender, err := aptos.NewAccountFromSigner(privateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		sender: sender,
	}, nil
}

func (a *Client) SendTransaction(payload aptos.TransactionPayload) (string, error) {
	rawTxn, err := a.client.BuildTransaction(a.sender.AccountAddress(), payload)
	if err != nil {
		return "", err
	}

	tx, err := a.client.SimulateTransaction(rawTxn, a.sender)
	if err != nil {
		return "", err
	}

	if !tx[0].Success {
		return "", errors.New("transaction failed")
	}

	signedTxn, err := rawTxn.SignedTransaction(a.sender)
	if err != nil {
		return "", err
	}

	submitResult, err := a.client.SubmitTransaction(signedTxn)
	if err != nil {
		return "", err
	}
	resp, err := a.client.WaitForTransaction(submitResult.Hash)
	if err != nil {
		fmt.Printf("Transaction %s is failure\n", submitResult.Hash)
		return "", err
	}

	if !resp.Success {
		return "", errors.New("transaction failed")
	}

	return submitResult.Hash, nil
}

func (a *Client) View(payload *aptos.ViewPayload) (any, error) {
	views, err := a.client.View(payload)
	if err != nil {
		return nil, err
	}
	if len(views) == 0 {
		return nil, errors.New("no view found")
	}

	return views[0], nil
}
