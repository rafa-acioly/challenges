package config

import "os"

// Config represents all resources that will be used
// by the application
type Config struct {
	DB    *DBSettings
}

// DBSettings define a basic config
// to access database
type DBSettings struct {
	Location          string
}

// New retrieves a set of configurations based on the env. variables
func New() Config {
	return Config{
		DB: &DBSettings{
			Location:          os.Getenv("DATABASE_LOCATION"),
		},
	}
}