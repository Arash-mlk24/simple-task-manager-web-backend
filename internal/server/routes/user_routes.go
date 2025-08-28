package routes

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(router *mux.Router, handler *handlers.UserHandler) {
	router.HandleFunc("/users", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/users", handler.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodGet)
}
