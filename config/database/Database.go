package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/w3tecch/go-api-boilerplate/config/env"
)

var gormConn *gorm.DB

// Connection returns gorm connection
func Connection() *gorm.DB {
	// Check if a connection allready exists
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	// Try to connect to the database
	conn, err := gorm.Open(env.Get().DBDialect, env.Get().DBConnection)
	if err != nil {
		logrus.Fatal("Could not connect to the database")
	}

	// Store the connection in package variable for furher request
	gormConn = conn

	// Disable logging
	gormConn.LogMode(false)

	return gormConn
}
