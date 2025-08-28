package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       int64  `json:"id" example:"1"`
	Username string `json:"username" example:"Arash_mlk24"`
	Email    string `json:"email" example:"arash.mros@gmail.com"`
}
