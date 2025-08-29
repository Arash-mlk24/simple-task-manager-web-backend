package seeders

import (
	"errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/auth"
	"gorm.io/gorm"
	"log"
)

func SeedRoles(db *gorm.DB) {
	roles := []entity.Role{
		{Title: auth.RoleAdmin},
		{Title: auth.RoleUser},
	}

	for _, role := range roles {
		var existing entity.Role
		err := db.Where("title = ?", role.Title).First(&existing).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&role).Error; err != nil {
					log.Printf("Failed to create role %s: %v", role.Title, err)
				} else {
					log.Printf("Seeded role: %s", role.Title)
				}
			}
		} else {
			log.Printf("Role (%v) already exists, skipped seeding.", role.Title)
		}
	}
}
