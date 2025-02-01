package app

import (
	"os"

	"gorm.io/gorm"
)

type AppConfig struct {
	Host   string
	Port   string
	InitDB bool
	DB     *gorm.DB
}

func (app *AppConfig) InitService() {
	app.Host = os.Getenv("APP_HOST")
	app.Port = os.Getenv("APP_PORT")

	// init db
	app.InitSqllite()
}
