package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id" example:"8420e01d-1a8a-471e-9c36-102163aed978"`
	Username string    `json:"username" example:"Arash_mlk24"`
	Email    string    `json:"email" example:"arash.mros@gmail.com"`
}
