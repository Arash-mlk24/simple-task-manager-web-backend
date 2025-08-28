package entity

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
