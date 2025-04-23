package blockchain

import (
	"github.com/aptos-labs/aptos-go-sdk"
)

type IClient interface {
	SendTransaction(payload aptos.TransactionPayload) (string, error)
	View(payload *aptos.ViewPayload) (any, error)
}
