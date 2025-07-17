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
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Подключаемся к базе данных
	var db *gorm.DB
	if cfg.IsSQLite() {
		db, err = gorm.Open(sqlite.Open(cfg.GetDatabaseDSN()), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(cfg.GetDatabaseDSN()), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автоматические миграции
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Создаем обработчики
	userHandler := handlers.NewUserHandler(db)

	// Настраиваем роутер
	router := gin.Default()

	// API роуты
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

	// Запускаем сервер
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
