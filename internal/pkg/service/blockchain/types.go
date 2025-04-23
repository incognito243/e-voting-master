package blockchain

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
)

type StringBCS string

func (s *StringBCS) MarshalBCS(ser *bcs.Serializer) {
	ser.WriteString(string(*s))
}

func (s *StringBCS) UnmarshalBCS(des *bcs.Deserializer) {
	*s = StringBCS(des.ReadString())
}

type AccountAddress aptos.AccountAddress

func (a *AccountAddress) MarshalBCS(ser *bcs.Serializer) {
	ser.FixedBytes(a[:])
}

func HexToAddress(hex string) aptos.AccountAddress {
	address := aptos.AccountAddress{}
	err := address.ParseStringRelaxed(hex)
	if err != nil {
		panic(fmt.Errorf("failed to parse address: %s", hex))
	}
	return address
}

type StringResult string

type Result struct {
	Name      StringResult `json:"name"`
	VoteCount int64        `json:"vote_count,string"`
}

func (r *StringResult) UnmarshalJSON(bytes []byte) error {
	hexStr := string(bytes)
	hexStr = strings.Trim(hexStr, `"`)
	if len(hexStr) >= 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return err
	}

	*r = StringResult(decoded)
	return nil
}
