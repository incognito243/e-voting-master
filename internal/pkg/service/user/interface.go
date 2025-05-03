package user

import (
	"context"

	"e-voting-mater/internal/pkg/entity"
)

type IService interface {
	CreateUser(ctx context.Context, user *entity.User, password string) error
	LoginUser(ctx context.Context, username string, password string, citizenId string) (*InfoUser, string, error)
	GetUserByUsername(ctx context.Context, username string) (*InfoUser, error)
	GetAllUsers(ctx context.Context) ([]*InfoUser, error)
	GetUserByCitizenID(ctx context.Context, citizenID string) (*InfoUser, error)

	VerifyUser(ctx context.Context, usernames []string, adminId, signatureHex string) error
	Vote(ctx context.Context, username string, serverId string, votingHex string, signatureHex string) error
}
