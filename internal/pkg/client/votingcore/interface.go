package votingcore

import "context"

type IClient interface {
	CreateVotingServer(ctx context.Context, serverConfig *CreateVotingServerRequest) (*CreateVotingServerResponse, error)
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	Vote(ctx context.Context, req *VoteRequest) error
	EndVote(ctx context.Context, req *EndVoteRequest) (*EndVoteResponse, error)
	PublishResult(ctx context.Context, serverId string) (*PublishResultResponse, error)
}
