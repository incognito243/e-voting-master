package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrackingRepo struct {
	db *gorm.DB
}

func NewTrackingRepo(db *gorm.DB) *TrackingRepo {
	return &TrackingRepo{
		db: db,
	}
}

func (t *TrackingRepo) Save(ctx context.Context, tracking *entity.Tracking) error {
	if err := t.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(tracking).Error; err != nil {
		return err
	}
	return nil
}

func (t *TrackingRepo) IsExist(ctx context.Context, username, serverId string) (bool, error) {
	var count int64
	if err := t.db.WithContext(ctx).Model(&entity.Tracking{}).Where("username = ? AND server_id = ?", username, serverId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
