package votingcore

import (
	"context"
	"fmt"
	"testing"

	"e-voting-mater/pkg/http"
)

func TestName(t *testing.T) {
	httpClient := http.NewClient(5, 5, 10)
	client := NewClient(httpClient, "http://localhost:8000", "123456")
	ctx := context.Background()

	server, err := client.CreateVotingServer(ctx, &CreateVotingServerRequest{
		NumberOfCandidates:    4,
		MaximumNumberOfVoters: 10,
		ServerName:            "test",
	})
	if err != nil {
		t.Fatalf("failed to create voting server: %v", err)
	}

	userId, err := client.CreateUser(ctx, &CreateUserRequest{
		Username: "username",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	err = client.Vote(ctx, &VoteRequest{
		UserID:      userId.UserID,
		ServerID:    server.ServerId,
		CandidateID: 1,
	})
	if err != nil {
		t.Fatalf("failed to vote: %v", err)
	}

	results, err := client.EndVote(ctx, &EndVoteRequest{
		ServerId: server.ServerId,
	})
	if err != nil {
		t.Fatalf("failed to open vote: %v", err)
	}

	fmt.Printf("Results: %v\n", results.Results)

	publishResult, err := client.PublishResult(ctx, server.ServerId)
	if err != nil {
		t.Fatalf("failed to publish publishResult: %v", err)
	}
	fmt.Printf("VoterPublicKey: %v\n", publishResult.VoterPublicKey)
	fmt.Printf("VoterVote: %v\n", publishResult.VoterVote)
	fmt.Printf("VoterSignedMessage: %v\n", publishResult.VoterSignedMessage)
	fmt.Printf("VoterProveOfWork: %v\n", publishResult.VoterProveOfWork)
	fmt.Printf("EncryptedPackage: %v\n", publishResult.EncryptedPackage)
	fmt.Printf("DecryptedPackage: %v\n", publishResult.DecryptedPackage)
	fmt.Printf("ResultPackage: %v\n", publishResult.ResultPackage)
	fmt.Printf("Results: %v\n", publishResult.Results)

}
