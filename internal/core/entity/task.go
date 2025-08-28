package entity

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	BaseEntity
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	DueDate      *time.Time `json:"due_date"`
	Tags         []Tag      `gorm:"many2many:task_tags;" json:"tags"`
	CollectionId uuid.UUID  `gorm:"type:uuid" json:"collection_id"`
	Collection   Collection `gorm:"foreignKey:CollectionId" json:"collection"`
}
