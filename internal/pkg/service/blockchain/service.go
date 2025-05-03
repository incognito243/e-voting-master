package blockchain

import (
	"encoding/json"

	"e-voting-mater/internal/pkg/client/blockchain"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
)

type Service struct {
	aptosClient blockchain.IClient
}

var service *Service

func NewService(aptosClient blockchain.IClient) *Service {
	if service == nil {
		service = &Service{
			aptosClient: aptosClient,
		}
	}
	return service
}

func Instance() *Service {
	return service
}

func (s *Service) Vote(hexAddress string, index uint64, contract string) (string, error) {
	address := HexToAddress(hexAddress)
	bytesAddress, err := bcs.Serialize(&address)
	if err != nil {
		return "", err
	}

	bytesIndex, err := bcs.SerializeU64(index)
	if err != nil {
		return "", err
	}

	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   module,
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

func (s *Service) AddCandidates(candidates []string, contract string) (string, error) {
	var bscCandidates []StringBCS
	for _, candidate := range candidates {
		bscCandidates = append(bscCandidates, StringBCS(candidate))
	}
	bytes, err := bcs.SerializeSequenceOnly(bscCandidates)
	if err != nil {
		return "", err
	}

	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   module,
			Function: FunctionAddCandidates,
			ArgTypes: nil,
			Args:     [][]byte{bytes},
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) EndVote(contract string) (string, error) {
	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   module,
			Function: FunctionEndVote,
			ArgTypes: nil,
			Args:     nil,
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) StartVote(contract string) (string, error) {
	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}
	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   module,
			Function: FunctionStartVote,
			ArgTypes: nil,
			Args:     nil,
		}})
	if err != nil {
		return "", err
	}
	return tx, err
}

func (s *Service) GiveRightToVote(hexAddress string, contract string) (string, error) {
	address := HexToAddress(hexAddress)
	bytesAddress, err := bcs.Serialize(&address)
	if err != nil {
		return "", err
	}

	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	tx, err := s.aptosClient.SendTransaction(aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module:   module,
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

func (s *Service) Admin(contract string) (string, error) {
	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   module,
		Function: ViewAdmin,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return "", err
	}

	return views.(string), nil
}

func (s *Service) GetResults(contract string) ([]*Result, error) {
	module, err := s.getModuleId(contract)
	if err != nil {
		return nil, err
	}

	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   module,
		Function: ViewGetResults,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return nil, err
	}

	bytesResults, err := json.Marshal(views)
	if err != nil {
		return nil, err
	}
	var results []*Result

	if err := json.Unmarshal(bytesResults, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Service) WinningCandidate(contract string) (string, error) {
	module, err := s.getModuleId(contract)
	if err != nil {
		return "", err
	}

	views, err := s.aptosClient.View(&aptos.ViewPayload{
		Module:   module,
		Function: ViewWinningCandidate,
		ArgTypes: nil,
		Args:     nil,
	})
	if err != nil {
		return "", err
	}

	bytesResults, err := json.Marshal(views)
	if err != nil {
		return "", err
	}
	var results StringResult

	if err := json.Unmarshal(bytesResults, &results); err != nil {
		return "", err
	}

	return string(results), nil
}

func (s *Service) getModuleId(hex string) (aptos.ModuleId, error) {
	address := aptos.AccountAddress{}
	err := address.ParseStringRelaxed(hex)
	if err != nil {
		return aptos.ModuleId{}, err
	}
	return aptos.ModuleId{
		Address: address,
		Name:    NameContract,
	}, nil
}
