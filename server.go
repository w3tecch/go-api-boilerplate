package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/w3tecch/go-api-boilerplate/app"
)

// ServerLog ...
// var ServerLog = config.Logger{Scope: "server"}

func main() {

	// Prepars server before running it
	app.Setup()

	// // Prints the version and the address of our api to the console
	// ServerLog.LineBreak()
	// ServerLog.WithFields(logrus.Fields{
	// 	"title":    config.Env().APITitle,
	// 	"version":  config.Env().APIVersion,
	// 	"env":      config.Env().Name,
	// 	"gin_mode": config.Env().GinMode,
	// }).Info("API information")
	// ServerLog.Info("Starting Server on http://localhost" + config.Env().APIPort)
	// ServerLog.LineBreak()

	// Listen on the given port on localhost
	app.Run()

}
