package entity

type User struct {
	BaseEntity
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
}

func (user *User) RolesAsString() []string {
	result := make([]string, len(user.Roles))
	for i, r := range user.Roles {
		result[i] = r.Title
	}
	return result
}
