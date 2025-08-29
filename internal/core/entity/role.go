package entity

type Role struct {
	BaseEntity
	Title string `gorm:"unique" json:"title"`
}
