package seeders

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"gorm.io/gorm"
	"log"
)

func SeedUsers(db *gorm.DB) {
	var adminRole entity.Role
	if err := db.First(&adminRole, "title = ?", "Admin").Error; err != nil {
		log.Printf("Admin role not found: %v", err)
		return
	}

	var existingUser entity.User
	if err := db.First(&existingUser, "email = ?", "admin@example.com").Error; err == nil {
		log.Println("Admin user already exists, skipped seeding.")
		return
	}

	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		log.Fatalf("An error occurred while hashing password: %v", err)
		return
	}

	user := entity.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: hashedPassword,
		Roles:    []entity.Role{adminRole},
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
		return
	}

	log.Println("Admin user seeded successfully")
}
