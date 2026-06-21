package repository

import (
	"context"
	"errors"
	"userfc/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository struct{
	Redis *redis.Client
	Database *gorm.DB
}


func NewUserRepository(redis *redis.Client, db *gorm.DB) *UserRepository {
	return &UserRepository{
		Redis: redis,
		Database: db,
	}
}

// GORM detects table automatically as changing type's name snake_case, and plural
// ex) User -> users


func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	// with timeout context
	err := r.Database.WithContext(ctx).Where("email = ?", email).Last(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// When no user is found, return nil.
			return nil, nil 
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetById(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	// with timeout context
	err := r.Database.WithContext(ctx).Where("id = ?", userID).Last(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// When no user is found, return nil.
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) (int64, error) {
	// with timeout context
	err := r.Database.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}