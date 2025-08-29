package repository

import (
	"context"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetByTitle(ctx context.Context, title string) (*entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (repository *roleRepository) GetByTitle(ctx context.Context, title string) (*entity.Role, error) {
	role := &entity.Role{}
	result := repository.db.WithContext(ctx).Where("title = ?", title).First(role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}
