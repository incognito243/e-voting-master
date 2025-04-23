package blockchain

import (
	"encoding/json"

	"e-voting-mater/internal/pkg/client/blockchain"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
)

type Service struct {
	aptosClient blockchain.IClient
	module      aptos.ModuleId
}

var service *Service

func NewService(aptosClient blockchain.IClient, hex string) *Service {
	if service == nil {
		address := aptos.AccountAddress{}
		err := address.ParseStringRelaxed(hex)
		if err != nil {
			panic(err)
		}
		service = &Service{
			aptosClient: aptosClient,
			module: aptos.ModuleId{
				// TODO: replace with configs
				Address: address,
				Name:    NameContract,
			},
		}
	}
	return service
}

func Instance() *Service {
	return service
}

func (s *Service) Vote(hexAddress string, index uint64) (string, error) {
	address := HexToAddress(hexAddress)
	bytesAddress, err := bcs.Serialize(&address)
	if err != nil {
		return "", err
	}

	bytesIndex, err := bcs.SerializeU64(index)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   s.module,
			Function: FunctionVote,
			ArgTypes: nil,
			Args: [][]byte{
				bytesAddress,
				bytesIndex,
			},
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) AddCandidates(candidates []string) (string, error) {
	var bscCandidates []StringBCS
	for _, candidate := range candidates {
		bscCandidates = append(bscCandidates, StringBCS(candidate))
	}
	bytes, err := bcs.SerializeSequenceOnly(bscCandidates)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   s.module,
			Function: FunctionAddCandidates,
			ArgTypes: nil,
			Args:     [][]byte{bytes},
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) EndVote() (string, error) {
	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   s.module,
			Function: FunctionEndVote,
			ArgTypes: nil,
			Args:     nil,
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) StartVote() (string, error) {
	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   s.module,
			Function: FunctionStartVote,
			ArgTypes: nil,
			Args:     nil,
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) GiveRightToVote(hexAddress string) (string, error) {
	address := HexToAddress(hexAddress)
	bytesAddress, err := bcs.Serialize(&address)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   s.module,
			Function: FunctionGiveRightToVote,
			ArgTypes: nil,
			Args: [][]byte{
				bytesAddress,
			},
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) Admin() (string, error) {
	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   s.module,
		Function: ViewAdmin,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return "", err
	}

	return views.(string), nil
}

func (s *Service) GetResults() ([]*Result, error) {
	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   s.module,
		Function: ViewGetResults,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return nil, err
	}

	bytesResults, err := json.Marshal(views)
	var results []*Result

	if err := json.Unmarshal(bytesResults, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Service) WinningCandidate() (string, error) {
	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   s.module,
		Function: ViewWinningCandidate,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return "", err
	}

	bytesResults, err := json.Marshal(views)
	var results StringResult

	if err := json.Unmarshal(bytesResults, &results); err != nil {
		return "", err
	}

	return string(results), nil
}
