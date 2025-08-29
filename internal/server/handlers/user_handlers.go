package handlers

import (
	"encoding/json"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service_errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
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
		utils.RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// GetUser
//
// @Summary Get user by id
// @Description Get a single user by their id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} dto.DocsApiResponse "ApiResponse wrapping UserResponse"
// @Failure 400 {object} dto.DocsApiResponse "Invalid UUID"
// @Failure 404 {object} dto.DocsApiResponse "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		serviceErr := service_errors.ErrIdNotValid
		response := utils.ApiFailure(serviceErr.Code, serviceErr.Message)
		utils.RespondJSON(w, serviceErr.HttpStatus, response)
		return
	}

	user, serviceErr := h.service.GetUser(r.Context(), id)
	if serviceErr != nil {
		response := utils.ApiFailure(serviceErr.Code, serviceErr.Message)
		utils.RespondJSON(w, serviceErr.HttpStatus, response)
		return
	}

	response := utils.ApiSuccess(user)
	utils.RespondJSON(w, http.StatusOK, response)
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
		utils.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, users)
}
