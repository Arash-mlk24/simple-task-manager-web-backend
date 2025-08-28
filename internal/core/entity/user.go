package entity

type User struct {
	BaseEntity
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
