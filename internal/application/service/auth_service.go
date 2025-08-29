package service

import (
	"context"
	"errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service_errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/infrastructure/repository"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/auth"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, *service_errors.ServiceError)
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return &authService{repository: repository}
}

func (service *authService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, *service_errors.ServiceError) {
	user, err := service.repository.GetByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &service_errors.ErrEmailOrPasswordMismatch
		}
		return nil, &service_errors.ErrInternal
	}

	if isPasswordIncorrect := !utils.CheckPassword(user.Password, request.Password); isPasswordIncorrect {
		return nil, &service_errors.ErrEmailOrPasswordMismatch
	}

	token, err := auth.GenerateJWT(user.Id.String(), user.RolesAsString())
	if err != nil {
		return nil, &service_errors.ErrInternal
	}

	return &dto.LoginResponse{AccessToken: token}, nil
}
