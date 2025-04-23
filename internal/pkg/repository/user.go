package repository

import (
	"context"

	"e-voting-mater/internal/pkg/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Save(ctx context.Context, user *entity.User) error {
	if err := u.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, userName string) (*entity.User, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) VerifyUsers(ctx context.Context, usernames []string) error {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.User{}).
		Where("username IN ?", usernames).
		Update("verified", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *UserRepo) GetUserByCitizenID(ctx context.Context, citizenId string) (*entity.User, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).Where("citizen_id = ?", citizenId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) SaveAdmin(ctx context.Context, user *entity.User) error {
	if err := u.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetAdmin(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := u.db.WithContext(ctx).Where("is_admin = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) GetAdminByAdminId(ctx context.Context, adminId string) (*entity.User, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).
		Where("citizen_id = ? AND is_admin = ?", adminId, true).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
