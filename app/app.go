package app

import (
	"github.com/joho/godotenv"
	"github.com/w3tecch/go-api-boilerplate/app/routes"
	"github.com/w3tecch/go-api-boilerplate/config/env"
	"github.com/w3tecch/go-api-boilerplate/config/logger"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
	"github.com/w3tecch/go-api-boilerplate/lib/seeder"
)

var appLog = logger.Logger{Scope: "app"}

func Setup() {
	LoadEnvConfig()
	SetupDatabase()
}

func Run() {
	// Listen on the given port on localhost
	r := routes.Router()
	r.Run(env.Get().APIPort)
}

func SetupDatabase() {
	// Migrates the database with out migrations scripts
	migrate.Sync()

	// Initialize seeder
	seeder.Sync()
}

func LoadEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		appLog.Fatal("Error loading .env file")
	}
}
