package main

import (
	"fmt"
	"log"

	"restapi/config"
	"restapi/handlers"
	"restapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Loading configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connecting to database
	var db *gorm.DB
	if cfg.IsSQLite() {
		db, err = gorm.Open(sqlite.Open(cfg.GetDatabaseDSN()), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(cfg.GetDatabaseDSN()), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automatic migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Creating handlers
	userHandler := handlers.NewUserHandler(db)

	// Setting up router
	router := gin.Default()

	// API routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// Starting server
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
