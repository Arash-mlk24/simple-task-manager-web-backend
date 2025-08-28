package handlers

import (
	"encoding/json"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
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

// Register
//
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User registration data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// GetUser
//
// @Summary Get user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// ListUsers
//
// @Summary List all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, users)
}
