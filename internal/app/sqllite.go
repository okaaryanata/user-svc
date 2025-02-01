package app

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/okaaryanata/user/internal/domain"
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

	defer c.migrateData()

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

func (c *AppConfig) migrateData() {
	isMigrateData, _ := strconv.ParseBool(os.Getenv("DB_MIGRATION"))
	if isMigrateData {
		var count int64
		c.DB.Model(&domain.User{}).Count(&count)

		if count == 0 {
			user := domain.User{Name: "Oka Aryanata", CreatedAt: time.Now().UnixMicro(), UpdatedAt: time.Now().UnixMicro()}
			c.DB.Create(&user)
			fmt.Println("User inserted:", user.Name)
		} else {
			fmt.Println("User already exists, no insertion.")
		}
	}
}
