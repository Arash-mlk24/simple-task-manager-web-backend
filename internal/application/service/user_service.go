package service

import (
	"context"
	"errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service_errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/infrastructure/repository"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUser(ctx context.Context, id uuid.UUID) (*dto.UserResponse, *service_errors.ServiceError)
	ListUsers(ctx context.Context) ([]dto.UserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (service *userService) Register(ctx context.Context, request dto.CreateUserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	savedUser, err := service.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{Id: savedUser.Id, Username: savedUser.Username, Email: savedUser.Email}, nil
}

func (service *userService) GetUser(ctx context.Context, id uuid.UUID) (*dto.UserResponse, *service_errors.ServiceError) {
	user, err := service.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &service_errors.ErrUserNotFound
		}

		return nil, &service_errors.ServiceError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Internal server error",
		}
	}

	return &dto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (service *userService) ListUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.UserResponse
	for _, u := range users {
		res = append(res, dto.UserResponse{Id: u.Id, Username: u.Username, Email: u.Email})
	}

	return res, nil
}
