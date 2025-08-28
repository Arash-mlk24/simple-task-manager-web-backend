package handlers

import (
	"encoding/json"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service service.UserService
}

func NewHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.service.Register(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// GetUser is a method for getting a user by its id
//
// handler function
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path   int  true  "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (handler *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	idStr := mux.Vars(request)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	user, err := handler.service.GetUser(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(writer).Encode(user)
}

func (handler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
