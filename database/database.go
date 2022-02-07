package database

import (
	"log"
	"time"
	"vts_api/database/migrations"
	"vts_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase () {
	
	dsn := "host=database user=postgres password=123456 dbname=vts_database port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// NowFunc: func() time.Time {
		// 	return time.Now().Local()
		// },
	})

	if err != nil {
		log.Fatal ("Database error: " + err.Error())
	}

	db = database
	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDabatase () *gorm.DB {
	return db
}

func Reset () {

	GetDabatase().Migrator().DropTable(&models.Alert{})
	GetDabatase().Migrator().DropTable(&models.Fleet{})
	GetDabatase().Migrator().DropTable(&models.Vehicle{})
	GetDabatase().Migrator().DropTable(&models.Position{})
	
	// Recreate
	migrations.RunMigrations(db)
}