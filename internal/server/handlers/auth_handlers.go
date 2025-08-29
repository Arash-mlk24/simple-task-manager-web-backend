package handlers

import (
	"encoding/json"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service_errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login
//
// @Summary Get token
// @Description Gets token for the user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "User login data"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errInvalidData := service_errors.ErrInvalidData
		response := utils.ApiFailure(errInvalidData.Code, errInvalidData.Message)
		utils.RespondJSON(w, errInvalidData.HttpStatus, response)
		return
	}

	loginResponse, serviceErr := h.service.Login(r.Context(), req)
	if serviceErr != nil {
		response := utils.ApiFailure(serviceErr.Code, serviceErr.Message)
		utils.RespondJSON(w, serviceErr.HttpStatus, response)
		return
	}

	response := utils.ApiSuccess(loginResponse)
	utils.RespondJSON(w, http.StatusOK, response)
}
