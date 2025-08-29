package server

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/infrastructure/repository"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/handlers"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/routes"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	userService := service.NewUserService(userRepository, roleRepository)
	authService := service.NewAuthService(userRepository)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterAuthRoutes(router, authHandler)

}
