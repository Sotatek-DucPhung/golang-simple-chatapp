package main

import (
	"gin-gorm-crud/internal/api"
	"gin-gorm-crud/internal/config"
	"gin-gorm-crud/internal/database"
	"gin-gorm-crud/logger"

	"github.com/gin-gonic/gin"
)

// @title Gin Gorm CRUD API
// @version 1.0
// @description This is a sample server for a CRUD application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log := logger.NewLogger()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Initialize database
	db, err := database.ConnectDatabase(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	// Initialize router
	router := gin.Default()
	// Setup routes
	api.SetupRoutes(router, db)

	// Start server
	err = router.Run(cfg.ServerPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
