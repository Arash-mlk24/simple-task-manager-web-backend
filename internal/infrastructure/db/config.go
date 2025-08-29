package db

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/infrastructure/seeders"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Config() *gorm.DB {
	dsn := "host=localhost user=postgres password=T@skM@nagerP@stgresDb dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate your entities
	err = db.AutoMigrate(
		&entity.Collection{},
		&entity.Tag{},
		&entity.Task{},
		&entity.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	seeders.SeedRoles(db)
	seeders.SeedUsers(db)

	return db
}
