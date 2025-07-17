package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	Path     string // для SQLite
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	// Пытаемся загрузить .env.dev для разработки, если не получается - обычный .env
	if err := godotenv.Load(".env.dev"); err != nil {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "postgres" // по умолчанию PostgreSQL
	}

	return &Config{
		Database: DatabaseConfig{
			Type:     dbType,
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Path:     os.Getenv("DB_PATH"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}, nil
}

func (c *Config) GetDatabaseDSN() string {
	if c.Database.Type == "sqlite" {
		return c.Database.Path
	}

	// PostgreSQL DSN
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Database.Host, c.Database.User, c.Database.Password,
		c.Database.Name, c.Database.Port, c.Database.SSLMode)
}

func (c *Config) IsSQLite() bool {
	return c.Database.Type == "sqlite"
}
