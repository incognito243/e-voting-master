package db

import (
	"time"

	"e-voting-mater/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres() (*gorm.DB, error) {
	gormDB, err := gorm.Open(
		postgres.Open(configs.PConn()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(configs.G.Log.Level)),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(configs.G.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configs.G.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.G.DB.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return gormDB, nil
}
