package votingserver

import (
	"context"

	"e-voting-mater/internal/pkg/entity"
)

type IService interface {
	CreateVotingServer(ctx context.Context, votingServer *entity.VotingServer, candidate []*Candidate, msgHex string) error
	GetVotingServerByID(ctx context.Context, serverId string) (*ExtendedVotingServer, error)
	GetAllVotingServers(ctx context.Context) ([]*ExtendedVotingServer, error)

	EndVote(ctx context.Context, adminId, serverId string, msgHex string) (*Candidate, map[string]int64, error)
	PublishVote(ctx context.Context, serverId string) (map[string]int64, error)
	ActiveVoting(ctx context.Context, adminId, serverId string, signatureHex string) error
}
