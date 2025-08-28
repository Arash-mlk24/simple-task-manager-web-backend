package repository

import (
	"context"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	result := repository.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil // Id and other DB-generated fields are already populated
}

func (repository *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	result := repository.db.WithContext(ctx).First(user, id) // SELECT * FROM users WHERE id = ?
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	result := repository.db.WithContext(ctx).Find(&users) // SELECT * FROM users
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
