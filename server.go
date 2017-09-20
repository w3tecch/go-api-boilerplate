package main

import (
	"os"

	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/w3tecch/go-api-boilerplate/app"
	"github.com/w3tecch/go-api-boilerplate/config"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
	"github.com/w3tecch/go-api-boilerplate/lib/seeder"
)

// ServerLog ...
var ServerLog = config.Logger{Scope: "server"}

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		ServerLog.Fatal("Error loading .env file")
	}

	// Migrates the database with out migrations scripts
	migrate.Sync()

	// Initialize seeder
	ServerLog.LineBreak()
	seeder.Sync()

	// Prints the version and the address of our api to the console
	ServerLog.LineBreak()
	ServerLog.WithFields(logrus.Fields{
		"title":    os.Getenv("API_TITLE"),
		"version":  os.Getenv("API_VERSION"),
		"env":      os.Getenv("ENVIRONMENT"),
		"gin_mode": os.Getenv("GIN_MODE"),
	}).Info("API information")
	ServerLog.Info("Starting Server on http://localhost" + os.Getenv("API_PORT"))
	ServerLog.LineBreak()

	// Listen on the given port on localhost
	r := app.Router()
	r.Run(os.Getenv("API_PORT"))

}
