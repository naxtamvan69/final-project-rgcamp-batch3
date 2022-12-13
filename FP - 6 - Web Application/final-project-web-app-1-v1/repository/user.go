package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	// creater varible user
	// get user data by id using gorm with context
	var user entity.User
	err := r.db.WithContext(ctx).Model(&entity.User{}).Select("*").Where("id = ?", id).Scan(&user).Error
	if err != nil {
		return entity.User{}, err
	} else if user.ID == 0 {
		return entity.User{}, nil
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	// create variable user
	// get user data by email using gorm with context
	var user entity.User
	err := r.db.WithContext(ctx).Model(&entity.User{}).Select("*").Where("email = ?", email).Scan(&user).Error
	if err != nil {
		return entity.User{}, err
	} else if user.ID == 0 {
		return entity.User{}, nil
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	// create user using gorm with context
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	// update user using gorm with context
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	// delete user by id using gorm with context
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
