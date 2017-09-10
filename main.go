package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/w3tecch/go-api-boilerplate/app"
	"github.com/w3tecch/go-api-boilerplate/app/config"
	"github.com/w3tecch/go-api-boilerplate/app/models"
)

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Prints the version and the address of our api to the console
	logrus.Info("Version is ", os.Getenv("API_VERSION"))
	logrus.Info("Starting Server on http://localhost:", os.Getenv("API_PORT"))

	// Creates the database schema
	migrateDatabase()

	// Listen on the given port on localhost
	r := app.Router()
	r.Run(os.Getenv("API_PORT"))
}

func migrateDatabase() {
	db := config.GetDatabaseConnection()

	// Migrate the given tables
	db.AutoMigrate(&models.User{})

}
