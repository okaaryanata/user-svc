package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/okaaryanata/user/internal/api/health"
	"github.com/okaaryanata/user/internal/api/middleware"
	"github.com/okaaryanata/user/internal/api/user"
	"github.com/okaaryanata/user/internal/app"
	"github.com/okaaryanata/user/internal/domain"
	"github.com/okaaryanata/user/internal/repository"
	"github.com/okaaryanata/user/internal/service"
)

func main() {
	startService()
}

func startService() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &app.AppConfig{}
	app.InitService()
	sqlDB, err := app.DB.DB()
	if err != nil {
		log.Fatal("failed to get raw SQL database object", err)
	}
	defer sqlDB.Close()

	// Repositories
	userRepository := repository.NewUserRepository(app.DB)

	// Services
	userSvc := service.NewUserService(userRepository)

	// Controllers
	healthController := health.NewHealthController()
	userController := user.NewUserController(userSvc)

	// Create main route
	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: middleware.GetListSkipLogPath(),
	}))
	router.Use(gin.Recovery())
	router.Use(middleware.SetCORSMiddleware())

	// Register main route
	mainRoute := router.Group(domain.MainRoute)

	// Register routes
	healthController.RegisterRoutes(mainRoute)
	userController.RegisterRoutes(mainRoute)

	host := fmt.Sprintf("%s:%s", app.Host, app.Port)
	router.Run(host)
}
