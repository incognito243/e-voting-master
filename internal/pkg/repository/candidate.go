package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"

	"gorm.io/gorm"
)

type CandidateRepo struct {
	db *gorm.DB
}

func NewCandidateRepo(db *gorm.DB) *CandidateRepo {
	return &CandidateRepo{db: db}
}

func (c *CandidateRepo) Save(ctx context.Context, candidate *entity.Candidate) error {
	if err := c.db.WithContext(ctx).
		Save(candidate).Error; err != nil {
		return err
	}
	return nil
}

func (c *CandidateRepo) GetByServerID(ctx context.Context, serverId string) ([]*entity.Candidate, error) {
	var candidates []*entity.Candidate
	if err := c.db.WithContext(ctx).
		Where("server_id = ?", serverId).
		Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

func (c *CandidateRepo) GetByIndex(ctx context.Context, serverId string, index int64) (*entity.Candidate, error) {
	var candidate entity.Candidate
	if err := c.db.WithContext(ctx).
		Where("server_id = ? AND candidate_index = ?", serverId, index).
		First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (c *CandidateRepo) GetByName(ctx context.Context, serverId string, name string) (*entity.Candidate, error) {
	var candidate entity.Candidate
	if err := c.db.WithContext(ctx).
		Where("server_id = ? AND candidate_name = ?", serverId, name).
		First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}
