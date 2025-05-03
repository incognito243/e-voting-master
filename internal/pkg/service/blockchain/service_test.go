package blockchain

import (
	"testing"

	"e-voting-mater/configs"
	"e-voting-mater/internal/pkg/client/blockchain"

	"github.com/aptos-labs/aptos-go-sdk"
)

func TestService(t *testing.T) {
	aptosClient, err := aptos.NewClient(aptos.NetworkConfig{
		Name:       "testnet",
		ChainId:    2,
		NodeUrl:    "https://fullnode.testnet.aptoslabs.com/v1",
		IndexerUrl: "https://api.testnet.aptoslabs.com/v1/graphql",
		FaucetUrl:  "",
	})
	if err != nil {
		t.Fatalf("failed to create aptos client: %v", err)
	}

	blockChainClient, _ := blockchain.NewClient(configs.VotingConfig{
		PrivateKeyName: "APTOS_PRIVATE_KEY",
	}, aptosClient)
	_ = NewService(blockChainClient)

	//tx, err := service.AddCandidates([]StringBCS{
	//	"simon",
	//	"james",
	//	"john",
	//})
	//if err != nil {
	//	t.Fatalf("failed to add candidates: %v", err)
	//}
	//t.Logf("Transaction hash: %s", tx)

	//voter := "0x3a2e759b988d2de0624c7779e84b1e38f15126462d82f0edc0e60d8ac85eb425"

	//tx, err := service.StartVote()
	//if err != nil {
	//	t.Fatalf("failed to start vote: %v", err)
	//}
	//t.Logf("Transaction hash: %s", tx)
	//
	//tx, err = service.GiveRightToVote(voter)
	//
	//tx, err = service.Vote(voter, 1)
	//if err != nil {
	//	t.Fatalf("failed to vote: %v", err)
	//}
	//t.Logf("Transaction hash: %s", tx)
	//
	//tx, err = service.EndVote()
	//if err != nil {
	//	t.Fatalf("failed to end vote: %v", err)
	//}
	//t.Logf("Transaction hash: %s", tx)

	//admin, err := service.Admin()
	//if err != nil {
	//	t.Fatalf("failed to get admin: %v", err)
	//}
	//t.Logf("Admin address: %s", admin)
	//
	//results, err := service.GetResults()
	//if err != nil {
	//	t.Fatalf("failed to get admin: %v", err)
	//}
	//t.Logf("Admin address: %+v", results[0])
	//t.Logf("Admin address: %+v", results[1])
	//t.Logf("Admin address: %+v", results[2])
	//
	//winningCandidate, err := service.WinningCandidate()
	//if err != nil {
	//	t.Fatalf("failed to get admin: %v", err)
	//}
	//t.Logf("Admin address: %s", winningCandidate)
}
