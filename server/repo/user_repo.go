package repo

import (
	"context"
	"log"
	"server/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateUser
func (repo *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	result := repo.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	var retrievedUser model.User
	result = repo.db.WithContext(ctx).First(&retrievedUser, "username = ?", user.Username)
	if result.Error != nil {
		return result.Error
	}

	log.Println("created user:", retrievedUser)
	return nil
}

// GetUser retrieves a user by ID or username
func (repo *UserRepo) GetUser(ctx context.Context, identifier interface{}, byID bool) (*model.User, error) {
	var user model.User
	var result *gorm.DB

	if byID {
		result = repo.db.WithContext(ctx).First(&user, identifier)
	} else {
		result = repo.db.WithContext(ctx).First(&user, "username = ?", identifier)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// UpdateUser updates a user by ID
func (repo *UserRepo) UpdateUser(ctx context.Context, identifier interface{}, byID bool, updatedUser *model.User) error {

	var result *gorm.DB
	if byID {
		result = repo.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", identifier).Updates(updatedUser)
	} else {
		result = repo.db.WithContext(ctx).Model(&model.User{}).Where("username= ?", identifier).Updates(updatedUser)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user by ID
func (repo *UserRepo) DeleteUser(ctx context.Context, userID uint) error {
	result := repo.db.WithContext(ctx).Delete(&model.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
