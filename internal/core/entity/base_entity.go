package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate => Hook: auto-generate UUID before insert if not set
func (base *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return
}
