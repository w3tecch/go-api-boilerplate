package main

import (
	"os"

	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/w3tecch/go-api-boilerplate/lib/logger"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
)

// ServerLog ...
var ServerLog = logger.Logger{Scope: "server"}

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		ServerLog.Fatal("Error loading .env file")
	}

	// Prints the version and the address of our api to the console
	ServerLog.WithFields(logrus.Fields{
		"title":    os.Getenv("API_TITLE"),
		"version":  os.Getenv("API_VERSION"),
		"env":      os.Getenv("ENVIRONMENT"),
		"gin_mode": os.Getenv("GIN_MODE"),
	}).Info("API information")
	ServerLog.Info("Starting Server on http://localhost" + os.Getenv("API_PORT"))

	// Creates the database schema
	migrate.Sync()

	// // Listen on the given port on localhost
	// r := app.Router()
	// r.Run(os.Getenv("API_PORT"))
}