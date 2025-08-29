package routes

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAuthRoutes(router *mux.Router, handler *handlers.AuthHandler) {
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
