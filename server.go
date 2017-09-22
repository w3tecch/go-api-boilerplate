package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"

	"github.com/w3tecch/go-api-boilerplate/app"
	"github.com/w3tecch/go-api-boilerplate/config/env"
	"github.com/w3tecch/go-api-boilerplate/config/logger"
)

// ServerLog ...
var ServerLog = logger.Logger{Scope: "server"}

func main() {

	// Prepars server before running it
	app.Setup()

	// Prints the version and the address of our api to the console
	ServerLog.LineBreak()
	ServerLog.WithFields(logrus.Fields{
		"title":    env.Get().APITitle,
		"version":  env.Get().APIVersion,
		"env":      env.Get().Name,
		"gin_mode": env.Get().GinMode,
	}).Info("API information")
	ServerLog.Info("Starting Server on http://localhost" + env.Get().APIPort)
	ServerLog.LineBreak()

	// Listen on the given port on localhost
	app.Run()

}
