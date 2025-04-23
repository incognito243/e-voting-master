package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"

	"gorm.io/gorm"
)

type VotingServerRepo struct {
	db *gorm.DB
}

func NewVotingServerRepo(db *gorm.DB) *VotingServerRepo {
	return &VotingServerRepo{
		db: db,
	}
}

func (v *VotingServerRepo) Save(ctx context.Context, votingServer *entity.VotingServer) error {
	if err := v.db.WithContext(ctx).Create(votingServer).Error; err != nil {
		return err
	}
	return nil
}

func (v *VotingServerRepo) GetByServerID(ctx context.Context, serverId string) (*entity.VotingServer, error) {
	var votingServer entity.VotingServer
	if err := v.db.WithContext(ctx).Where("server_id = ?", serverId).First(&votingServer).Error; err != nil {
		return nil, err
	}
	return &votingServer, nil
}

func (v *VotingServerRepo) GetAll(ctx context.Context) ([]*entity.VotingServer, error) {
	var votingServers []*entity.VotingServer
	if err := v.db.WithContext(ctx).Find(&votingServers).Error; err != nil {
		return nil, err
	}
	return votingServers, nil
}

func (v *VotingServerRepo) GetByAdminId(ctx context.Context, adminId string) ([]*entity.VotingServer, error) {
	var votingServers []*entity.VotingServer
	if err := v.db.WithContext(ctx).Where("admin_id = ?", adminId).Find(&votingServers).Error; err != nil {
		return nil, err
	}
	return votingServers, nil
}

func (v *VotingServerRepo) OpenVote(ctx context.Context, serverId string, results string) error {
	err := v.db.WithContext(ctx).
		Model(&entity.VotingServer{}).
		Where("server_id = ?", serverId).
		Updates(map[string]interface{}{
			"opened_vote": true,
			"results":     results,
		}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (v *VotingServerRepo) ActiveServer(ctx context.Context, serverId string) error {
	err := v.db.WithContext(ctx).
		Model(&entity.VotingServer{}).
		Where("server_id = ?", serverId).
		Update("active", true).Error
	if err != nil {
		return err
	}
	return nil

}
