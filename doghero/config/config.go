package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents all resources that will be used
// by the application
type Config struct {
	DB *DBSettings
}

// DBSettings define a basic config
// to access database
type DBSettings struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

// New retrieves a set of configurations based on the env. variables
func New() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DB: &DBSettings{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Name:     os.Getenv("DATABASE_NAME"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
	}
}
