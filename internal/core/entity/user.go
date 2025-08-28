package entity

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
