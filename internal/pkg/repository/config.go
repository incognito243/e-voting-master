package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"

	"gorm.io/gorm"
)

type ConfigRepo struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) *ConfigRepo {
	return &ConfigRepo{db: db}
}

func (r *ConfigRepo) Save(ctx context.Context, key, value string) error {
	err := r.db.WithContext(ctx).Model(&entity.Configs{}).Where("key = ?", key).Update("value", value).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ConfigRepo) Get(ctx context.Context, key string) (string, error) {
	var config entity.Configs
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&config).Error
	if err != nil {
		return "", err
	}
	return config.Value, nil
}
