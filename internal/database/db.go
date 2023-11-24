package database

import (
	"fmt"
	"golang/config"
	"golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	// dsn := "host=postgres user=postgres password=admin dbname=jportal port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DatabaseConfig.DB_Host, cfg.DatabaseConfig.DB_User, cfg.DatabaseConfig.DB_Pswd, cfg.DatabaseConfig.DB_Name, cfg.DatabaseConfig.DB_Port, cfg.DatabaseConfig.DB_Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.Company{}, &models.Job{}, &models.Location{}, &models.Technology{}, &models.Qualification{}, &models.Shift{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}

	return db, nil

}
