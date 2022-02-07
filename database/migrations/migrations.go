package migrations

import (
	"vts_api/models"

	"gorm.io/gorm"
)

func RunMigrations (db *gorm.DB) {
	db.AutoMigrate(models.Alert{})
	db.AutoMigrate(models.Fleet{})
	db.AutoMigrate(models.Vehicle{})
	db.AutoMigrate(models.Position{})
}