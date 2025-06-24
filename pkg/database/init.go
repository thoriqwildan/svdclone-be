package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/thoriqwildan/svdclone-be/pkg/config"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "password"),
		config.GetEnv("DB_NAME", "postgres"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_SSLMODE", "disable"),
		config.GetEnv("DB_TIMEZONE", "Asia/Jakarta"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Info("Database connection established successfully")

	// Ini buat auto migrate coy
	err = DB.AutoMigrate(&models.User{}, &models.PaymentMethod{}, &models.PaymentChannel{})
	if err != nil {
		log.Fatal("Failed to auto migrate models:", err)
	} else {
		log.Info("Database models migrated successfully")
	}
}