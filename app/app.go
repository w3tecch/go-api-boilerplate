package app

import (
	"github.com/w3tecch/go-api-boilerplate/app/routes"
	"github.com/w3tecch/go-api-boilerplate/config/env"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
	"github.com/w3tecch/go-api-boilerplate/lib/seeder"
)

func Setup() {
	env.LoadEnvConfig()
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
