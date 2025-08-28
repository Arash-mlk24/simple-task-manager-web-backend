package server

import (
	"database/sql"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/infrastructure/repository"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/handlers"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/routes"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewHandler(userService)

	routes.RegisterUserRoutes(router, userHandler)
}
