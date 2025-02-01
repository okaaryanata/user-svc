package app

import (
	"errors"
	"os"

	"github.com/okaaryanata/user/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (c *AppConfig) InitSqllite() error {
	dbFile := os.Getenv("DB_FILE")
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			return errors.New("failed to create database file")
		}
		file.Close()
	}

	// Initialize Database
	var err error
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect to database")
	}
	db.AutoMigrate(&domain.User{})

	c.DB = db

	return nil
}
