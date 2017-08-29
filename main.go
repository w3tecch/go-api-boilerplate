package main

import (
	"log"
	"net/http"
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

	// Check database connection
	db := config.GetDatabaseConnection()
	defer db.Close()

	// Prints the version and the address of our api to the console
	logrus.Info("Version is ", os.Getenv("API_VERSION"))
	logrus.Info("Starting Server on http://localhost:", os.Getenv("API_PORT"))

	// // Set log level
	// switch os.Getenv("LOG_LEVEL") {
	// case "debug":
	// 	logrus.SetLevel(logrus.ErrorLevel)
	// case "info":
	// 	logrus.SetLevel(logrus.ErrorLevel)
	// case "warn":
	// 	logrus.SetLevel(logrus.ErrorLevel)
	// default:
	// 	logrus.SetLevel(logrus.ErrorLevel)
	// }

	// Creates the database schema
	migrateDatabase()

	// Server router on given port and attach the cors headers
	server := app.NewServer()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), server))
}

func migrateDatabase() {
	db := config.GetDatabaseConnection()

	// Migrate the given tables
	db.AutoMigrate(&models.User{})

}
