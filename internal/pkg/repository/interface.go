package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"
)

type IUserRepo interface {
	Save(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	VerifyUsers(ctx context.Context, username []string) error
	GetUserByCitizenID(ctx context.Context, citizenId string) (*entity.User, error)
	SaveAdmin(ctx context.Context, user *entity.User) error
	GetAdmin(ctx context.Context) ([]*entity.User, error)
	GetAdminByAdminId(ctx context.Context, adminId string) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
}

type IVotingServerRepo interface {
	Save(ctx context.Context, votingServer *entity.VotingServer) error
	GetByServerID(ctx context.Context, serverId string) (*entity.VotingServer, error)
	GetAll(ctx context.Context) ([]*entity.VotingServer, error)
	GetByAdminId(ctx context.Context, adminId string) ([]*entity.VotingServer, error)
	OpenVote(ctx context.Context, serverId string, result string) (*entity.VotingServer, error)
	ActiveServer(ctx context.Context, serverName string) error
}

type ICandidateRepo interface {
	Save(ctx context.Context, candidate *entity.Candidate) error
	GetByServerID(ctx context.Context, serverId string) ([]*entity.Candidate, error)
	GetByIndex(ctx context.Context, serverId string, index int64) (*entity.Candidate, error)
	GetByName(ctx context.Context, serverId string, name string) (*entity.Candidate, error)
}

type IConfigRepo interface {
	Save(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}
