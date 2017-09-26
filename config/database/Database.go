package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/w3tecch/go-api-boilerplate/config/env"
	"github.com/w3tecch/go-api-boilerplate/config/logger"
	"github.com/w3tecch/go-api-boilerplate/constants"
)

var gormConn *gorm.DB
var dbLog = logger.Logger{Scope: "config.database"}

// Connection returns gorm connection
func Connection() *gorm.DB {
	// Check if a connection allready exists
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	// Try to connect to the database
	conn, err := gorm.Open(env.Get().DBDialect, env.Get().DBConnection)
	if err != nil {
		dbLog.Fatal("Could not connect to the database")
	}

	// Store the connection in package variable for furher request
	gormConn = conn

	// Disable logging
	gormConn.LogMode(false)

	return gormConn
}

func Clear() {
	db := Connection()
	db.DropTableIfExists(constants.TableUsers)
	db.DropTableIfExists(constants.TableSeeder)
	db.DropTableIfExists(constants.TableVersion)
}
