package env

import (
	"os"
)

type EnvConfig struct {
	Name         string
	GinMode      string
	APITitle     string
	APIVersion   string
	APIPort      string
	DBDialect    string
	DBConnection string
	LogLevel     string
	Auth0BaseURL string
}

func Get() *EnvConfig {
	return &EnvConfig{
		Name:         os.Getenv("ENVIRONMENT"),
		GinMode:      os.Getenv("GIN_MODE"),
		APITitle:     os.Getenv("API_TITLE"),
		APIVersion:   os.Getenv("API_VERSION"),
		APIPort:      os.Getenv("API_PORT"),
		DBDialect:    os.Getenv("DB_DIALECT"),
		DBConnection: os.Getenv("DB_CONNECTION"),
		LogLevel:     os.Getenv("LOG_LEVEL"),
		Auth0BaseURL: os.Getenv("AUTH0_BASE_URL"),
	}
}
